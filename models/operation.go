package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"tezosign/types"
)

type Contract struct {
	ID                      uint64        `gorm:"column:ctr_id;primaryKey"`
	Address                 types.Address `gorm:"column:ctr_address"`
	LastOperationBlockLevel uint64        `gorm:"column:ctr_last_block_level"`
}

type RequestStatus string

const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusRejected = "rejected"
	//Status for incoming transfers
	StatusSuccess = "success"
)

type Request struct {
	ID         uint64                   `gorm:"column:req_id;primaryKey" json:"-"`
	Hash       string                   `gorm:"column:req_hash" json:"operation_id,omitempty"`
	ContractID uint64                   `gorm:"column:ctr_id" json:"-"`
	Counter    *int64                   `gorm:"column:req_counter" json:"nonce,omitempty"`
	Status     RequestStatus            `gorm:"column:req_status;default:pending" json:"status"`
	CreatedAt  types.JSONTimestamp      `gorm:"column:req_created_at" json:"created_at"`
	Info       ContractOperationRequest `gorm:"column:req_info" json:"operation_info"`
	NetworkID  string                   `gorm:"column:req_network_id" json:"network_id"`

	OperationID *string `gorm:"column:req_operation_id" json:"tx_id,omitempty"`

	//Previous state of storage
	StorageDiff *StorageDiff `gorm:"column:req_storage_diff" json:"storage_diff,omitempty"`

	//Internal operation nonce
	Nonce sql.NullInt64 `gorm:"column:req_nonce" json:"-"`
}

type StorageDiff struct {
	Counter   Diff `json:"counter"`
	Threshold Diff `json:"threshold"`
	Keys      Diff `json:"keys"`
}

func (r *StorageDiff) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	data, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid type")
	}

	if len(data) == 0 {
		return nil
	}

	err = json.Unmarshal([]byte(data), r)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %s", err.Error())
	}

	return nil
}

func (r StorageDiff) Value() (driver.Value, error) {

	bt, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return string(bt), nil
}

type Diff struct {
	Previous interface{} `json:"previous,omitempty"`
	Current  interface{} `json:"current"`
}

type RequestReport struct {
	Request
	Signatures Signatures `gorm:"column:signatures" json:"signatures,omitempty"`
}

type OperationToSignResp struct {
	OperationID string        `json:"operation_id"`
	Payload     types.Payload `json:"payload"`
	PayloadJSON string        `json:"payload_json"`
}

type OperationSignatureResp struct {
	SigCount  int64 `json:"sig_count"`
	Threshold int64 `json:"threshold"`
}

type OperationParameter struct {
	Entrypoint string `json:"entrypoint"`
	Value      string `json:"value,omitempty"`
}

//Indexer Tezos operation
type TezosOperation struct {
	Level         uint64              `gorm:"column:Level"`
	Timestamp     types.JSONTimestamp `gorm:"column:Timestamp"`
	OpHash        string              `gorm:"column:OpHash"`
	BakerFee      uint64              `gorm:"column:BakerFee"`
	StorageFee    uint64              `gorm:"column:StorageFee"`
	AllocationFee uint64              `gorm:"column:AllocationFee"`
	Status        int                 `gorm:"column:Status"`
	Errors        string              `gorm:"column:Errors"`
}

type TransactionOperation struct {
	TezosOperation

	Amount     uint64        `gorm:"column:Amount"`
	Entrypoint string        `gorm:"column:Entrypoint"`
	Nonce      sql.NullInt64 `gorm:"column:Nonce"`
	//JsonParameters     []byte `gorm:"column:JsonParameters"`
	RawParameters      *types.TZKTPrim `gorm:"column:RawParameters"`
	InternalOperations uint64          `gorm:"column:InternalOperations"`
}

type RevealOperation struct {
	TezosOperation
	PublicKey types.PubKey `gorm:"column:PublicKey"`
}

type OriginationOperation struct {
	TezosOperation
	ContractAddress types.Address `gorm:"column:Address"`
}
