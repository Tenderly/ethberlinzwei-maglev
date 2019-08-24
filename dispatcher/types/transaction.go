package types

type Reference struct {
	VHashId       string `json:"v_hash_id" bson:"v_hash_id"`
	VHashPosition uint64 `json:"v_hash_position" bson:"v_hash_position"`
}

type Transaction struct {
	VHash  string `json:"v_hash" bson:"v_hash"`
	VNonce uint64 `json:"v_nonce" bson:"v_nonce"`

	FromTime uint64 `json:"from_time" bson:"from_time"`
	ToTime   uint64 `json:"to_time" bson:"to_time"`

	OperationType uint8  `json:"operation_type" bson:"operation_type"`
	Identity      string `json:"identity" bson:"identity"`

	To      string `json:"to" bson:"to"`
	Value   uint64 `json:"value" bson:"value"`
	Gas     uint64 `json:"gas" bson:"gas"`
	Payload string `json:"payload" bson:"payload"`

	TxHash      string `json:"tx_hash" bson:"tx_hash"`
	BlockNumber uint64 `json:"block_number" bson:"block_number"`
	GasPrice    uint64 `json:"gas_price" bson:"gas_price"`
	Success     bool   `json:"success" bson:"success"`
}

type TransactionService interface {
	Insert(tx *Transaction) error
	Find() ([]*Transaction, error)
}
