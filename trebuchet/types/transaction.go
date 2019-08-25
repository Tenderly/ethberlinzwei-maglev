package types

type Reference struct {
	VHashId       string `json:"v_hash_id" bson:"v_hash_id"`
	VHashPosition int64  `json:"v_hash_position" bson:"v_hash_position"`
}

type Transaction struct {
	VHash  string `json:"v_hash" bson:"v_hash"`
	VNonce int64  `json:"v_nonce" bson:"v_nonce"`

	FromTime int64 `json:"from_time" bson:"from_time"`
	ToTime   int64 `json:"to_time" bson:"to_time"`

	OperationType int64  `json:"operation_type" bson:"operation_type"`
	Identity      string `json:"identity" bson:"identity"`

	To       string `json:"to" bson:"to"`
	Value    int64  `json:"value" bson:"value"`
	Gas      int64  `json:"gas" bson:"gas"`
	Payload  string `json:"payload" bson:"payload"`
	SendTime int64  `json:"send_time" bson:"send_time"`

	TxHash   string `json:"tx_hash" bson:"tx_hash"`
	GasPrice int64  `json:"gas_price" bson:"gas_price"`

	EstimateGas int64 `json:"estimate_gas" bson:"estimate_gas"`
	SpentGas    int64 `json:"spent_gas" bson:"spent_gas"`

	MinedTime   int64 `json:"mined_time" bson:"mined_time"`
	BlockNumber int64 `json:"block_number" bson:"block_number"`
	Success     bool  `json:"success" bson:"success"`
}

type TransactionService interface {
	Insert(*Transaction) error
	InsertRaw(interface{}) error
	InsertMany([]interface{}) error
	Find() ([]*Transaction, error)
	FindUnsent(now int64) ([]*Transaction, error)
	UpdateTxsHashGasPriceAndTime(vHashes []string, txHash string, gasPrice, now int64) error
	UpdateTxsBlockNumberGasPriceAndTime(vHashes []string, blockNumber, gasPrice, now int64) error
}
