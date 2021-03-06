package contract

import (
	"encoding/hex"
	"math/big"
	"tezosign/models"
	"tezosign/types"

	"blockwatch.cc/tzindex/micheline"
	"github.com/anchorageoss/tezosprotocol/v2"
)

const (
	TextWatermark  = 0x05
	MainEntrypoint = "main_parameter"
	emptyOperation = `[{"prim":"DROP"},{"args":[{"prim":"operation"}],"prim":"NIL"}]`
)

type Operation struct {
	Entrypoint string          `json:"entrypoint"`
	Value      *micheline.Prim `json:"value"`
}

func BuildContractSignPayload(networkID string, counter int64, operationParams models.ContractOperationRequest) (resp types.Payload, jsonResp string, err error) {

	networkArgs, err := buildNetworkMichelsonArgs(networkID, operationParams.ContractID)
	if err != nil {
		return resp, jsonResp, err
	}

	params, err := buildActionMichelsonArgs(counter, operationParams)
	if err != nil {
		return resp, jsonResp, err
	}

	operation := &micheline.Prim{
		Type:   micheline.PrimBinary,
		OpCode: micheline.D_PAIR,
		Args:   []*micheline.Prim{networkArgs, params},
	}

	bt, err := operation.MarshalBinary()
	if err != nil {
		return resp, jsonResp, err
	}

	operationJson, err := operation.MarshalJSON()
	if err != nil {
		return resp, jsonResp, err
	}

	bt = append([]byte{TextWatermark}, bt...)

	return types.Payload(hex.EncodeToString(bt)), string(operationJson), nil
}

func BuildRejectSignPayload(networkID string, counter int64, contractAddress types.Address) (resp types.Payload, respJson string, err error) {
	return BuildContractSignPayload(networkID, counter, models.ContractOperationRequest{
		ContractID:    contractAddress,
		Type:          models.CustomPayload,
		CustomPayload: emptyOperation,
	})
}

//Network and contract address params
func buildNetworkMichelsonArgs(networkID string, contractID types.Address) (params *micheline.Prim, err error) {
	_, encodedNetworkID, err := tezosprotocol.Base58CheckDecode(networkID)
	if err != nil {
		return params, err
	}

	encodedContractAddress, err := contractID.MarshalBinary()
	if err != nil {
		return params, err
	}

	//Network and contract to bytes
	networkArgs := &micheline.Prim{
		Type:   micheline.PrimBinary,
		OpCode: micheline.D_PAIR,
		Args: []*micheline.Prim{
			{
				Type:   micheline.PrimBytes,
				OpCode: micheline.T_BYTES,
				Bytes:  encodedNetworkID,
			},
			{
				Type:   micheline.PrimBytes,
				OpCode: micheline.T_BYTES,
				Bytes:  encodedContractAddress,
			},
		},
	}

	return networkArgs, nil
}

func buildActionMichelsonArgs(counter int64, operationParams models.ContractOperationRequest) (params *micheline.Prim, err error) {

	actionArgs, err := buildActionCallMichelsonArgs(operationParams)
	if err != nil {
		return params, err
	}

	//Init params
	params = &micheline.Prim{
		Type:   micheline.PrimBinary,
		OpCode: micheline.D_PAIR,
		Args: []*micheline.Prim{
			//Add counter as left pair value
			{
				Type:   micheline.PrimInt,
				OpCode: micheline.T_INT,
				Int:    big.NewInt(counter),
			},
			//Right pair elem action args
			actionArgs,
		},
	}

	return params, nil
}

func buildActionCallMichelsonArgs(operationParams models.ContractOperationRequest) (params *micheline.Prim, err error) {

	actionParams, err := buildActionParams(operationParams)
	if err != nil {
		return params, err
	}

	//Build path
	path, err := buildMichelsonPath(operationParams.Type, actionParams)
	if err != nil {
		return params, err
	}

	return path, nil
}
