package services

import (
	"database/sql"
	"tezosign/common/apperrors"
	"tezosign/models"
	"tezosign/services/contract"
	"tezosign/types"

	"blockwatch.cc/tzindex/micheline"
)

func (s *ServiceFacade) BuildVestingContractInitStorage(storageReq models.VestingContractStorageRequest) (resp []byte, err error) {

	resp, err = contract.BuildVestingContractStorage(storageReq.VestingAddress, storageReq.DelegateAdmin, storageReq.Timestamp, storageReq.SecondsPerTick, storageReq.TokensPerTick)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *ServiceFacade) VestingContractInfo(contractID types.Address) (info models.VestingContractInfo, err error) {

	indexerRepo := s.indexerRepoProvider.GetIndexer()

	account, isFound, err := indexerRepo.GetAccount(contractID)
	if err != nil {
		return info, err
	}

	if !isFound {
		return info, apperrors.New(apperrors.ErrNotFound, "contract")
	}

	var delegate types.Address
	if account.DelegateID.Valid {
		delegateAccount, isFound, err := indexerRepo.GetAccountByID(uint64(account.DelegateID.Int64))
		if err != nil {
			return info, err
		}

		if !isFound {
			return info, apperrors.New(apperrors.ErrNotFound, "delegate")
		}

		delegate = delegateAccount.Address
	}

	script, storage, err := s.getContractScriptAndStorage(contractID)
	if err != nil {
		return info, err
	}

	storageContainer, err := contract.NewVestingContractStorageContainer(micheline.Script{
		Code: &micheline.Code{
			Param:   script.ParameterSchema.MichelinePrim(),
			Storage: script.StorageSchema.MichelinePrim(),
			Code:    script.CodeSchema.MichelinePrim(),
		},
		Storage: storage.RawValue.MichelinePrim(),
	})
	if err != nil {
		return info, err
	}

	openedTicks := storageContainer.OpenedTicks()

	//Mul to tokens per tick
	openedAmount := openedTicks * storageContainer.TokensPerTick
	vestedAmount := storageContainer.VestedTicks * storageContainer.TokensPerTick

	return models.VestingContractInfo{
		Balance:       account.Balance,
		OpenedBalance: openedAmount,
		Delegate:      delegate,
		Storage: models.VestingContractStorageRequest{
			VestingAddress: storageContainer.VestingAddress,
			DelegateAdmin:  storageContainer.DelegateAdmin,
			VestedAmount:   vestedAmount,
			Timestamp:      storageContainer.Timestamp,
			SecondsPerTick: storageContainer.SecondsPerTick,
			TokensPerTick:  storageContainer.TokensPerTick,
		},
	}, nil
}

func (s *ServiceFacade) VestingContractOperation(req models.VestingContractOperation) (param models.OperationParameter, err error) {

	value, entrypoint, err := contract.VestingContractParamAndEntrypoint(req)
	if err != nil {
		return param, err
	}

	return models.OperationParameter{
		Entrypoint: entrypoint,
		Value:      string(value),
	}, nil
}

func (s *ServiceFacade) VestingsList(userPubKey types.PubKey, contractAddress types.Address) (vestings []models.Vesting, err error) {

	contract, isFound, err := s.repoProvider.GetContract().GetContract(contractAddress)
	if err != nil {
		return vestings, err
	}

	if !isFound {
		return vestings, apperrors.New(apperrors.ErrNotFound, "contract")
	}

	isOwner, err := s.GetUserAllowance(userPubKey, contractAddress)
	if err != nil {
		return vestings, err
	}

	//For viewer return empty arr
	if !isOwner {
		return vestings, nil
	}

	vestings, err = s.repoProvider.GetVesting().GetVestingsList(contract.ID)
	if err != nil {
		return vestings, err
	}

	indexerRepo := s.indexerRepoProvider.GetIndexer()
	for i := range vestings {
		acc, _, err := indexerRepo.GetAccount(vestings[i].Address)
		if err != nil {
			return vestings, err
		}

		vestings[i].Balance = acc.Balance
	}

	return vestings, err
}

func (s *ServiceFacade) ContractVesting(userPubKey types.PubKey, contractAddress types.Address, reqVesting models.Vesting) (vesting models.Vesting, err error) {

	contract, isFound, err := s.repoProvider.GetContract().GetContract(contractAddress)
	if err != nil {
		return vesting, err
	}

	if !isFound {
		return vesting, apperrors.New(apperrors.ErrNotFound, "contract")
	}

	//Сheck contract for vesting type
	isVestingContract, err := s.checkVestingContract(reqVesting.Address)
	if err != nil {
		return vesting, err
	}

	if !isVestingContract {
		return vesting, apperrors.New(apperrors.ErrBadParam, "not vesting contract")
	}

	vestingRepo := s.repoProvider.GetVesting()
	vesting, isFound, err = vestingRepo.GetVesting(contract.ID, reqVesting.Address)
	if err != nil {
		return vesting, err
	}

	//Already created
	if isFound {
		return vesting, apperrors.New(apperrors.ErrAlreadyExists, "asset")
	}

	reqVesting.ContractID = sql.NullInt64{
		Int64: int64(contract.ID),
		Valid: true,
	}

	vestingAcc, _, err := s.indexerRepoProvider.GetIndexer().GetAccount(reqVesting.Address)
	if err != nil {
		return vesting, err
	}

	reqVesting.Balance = vestingAcc.Balance

	err = vestingRepo.CreateVesting(reqVesting)
	if err != nil {
		return vesting, err
	}

	return reqVesting, nil
}

func (s *ServiceFacade) ContractVestingEdit(userPubKey types.PubKey, contractAddress types.Address, reqVesting models.Vesting) (vesting models.Vesting, err error) {

	contract, isFound, err := s.repoProvider.GetContract().GetContract(contractAddress)
	if err != nil {
		return vesting, err
	}

	if !isFound {
		return vesting, apperrors.New(apperrors.ErrNotFound, "contract")
	}

	vestingRepo := s.repoProvider.GetVesting()
	vesting, isFound, err = vestingRepo.GetVesting(contract.ID, reqVesting.Address)
	if err != nil {
		return vesting, err
	}

	//Not created
	if !isFound {
		return vesting, apperrors.New(apperrors.ErrNotFound, "asset")
	}

	//Global asset cannot be edited
	if !vesting.ContractID.Valid {
		return vesting, apperrors.New(apperrors.ErrNotAllowed, "global asset")
	}

	vesting.Name = reqVesting.Name

	vestingAcc, _, err := s.indexerRepoProvider.GetIndexer().GetAccount(vesting.Address)
	if err != nil {
		return vesting, err
	}

	vesting.Balance = vestingAcc.Balance

	err = vestingRepo.UpdateVesting(vesting)
	if err != nil {
		return vesting, err
	}

	return vesting, nil
}

func (s *ServiceFacade) RemoveContractVesting(userPubKey types.PubKey, contractAddress types.Address, vesting models.Vesting) (err error) {

	contract, isFound, err := s.repoProvider.GetContract().GetContract(contractAddress)
	if err != nil {
		return err
	}

	if !isFound {
		return apperrors.New(apperrors.ErrNotFound, "contract")
	}

	vestingRepo := s.repoProvider.GetVesting()
	vesting, isFound, err = vestingRepo.GetVesting(contract.ID, vesting.Address)
	if err != nil {
		return err
	}

	if !isFound {
		return apperrors.New(apperrors.ErrNotFound, "asset")
	}

	//Global asset cannot be removed
	if !vesting.ContractID.Valid {
		return apperrors.New(apperrors.ErrNotAllowed, "global asset")
	}

	err = vestingRepo.DeleteContractVesting(vesting.ID)
	if err != nil {
		return err
	}

	return nil
}
