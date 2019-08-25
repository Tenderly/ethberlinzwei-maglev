package mongo

import (
	"context"

	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/types"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

const Collection = "transactions"

type TransactionService struct {
	client *Client
	ctx    context.Context
}

func NewTransactionService(client *Client) *TransactionService {
	return &TransactionService{
		client: client,
	}
}

func (ts *TransactionService) Insert(tx *types.Transaction) error {
	_, err := ts.client.DB.Collection(Collection).InsertOne(ts.ctx, tx)

	return errors.Wrap(err, "insert transaction")
}

func (ts *TransactionService) InsertRaw(tx interface{}) error {
	_, err := ts.client.DB.Collection(Collection).InsertOne(ts.ctx, tx)

	return errors.Wrap(err, "insert transaction")
}

func (ts *TransactionService) InsertMany(txs []interface{}) error {
	_, err := ts.client.DB.Collection(Collection).InsertMany(ts.ctx, txs)

	return errors.Wrap(err, "insert transaction")
}

func (ts *TransactionService) Find() ([]*types.Transaction, error) {
	cursor, err := ts.client.DB.Collection(Collection).Find(ts.ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "get transactions")
	}

	var transactions []*types.Transaction
	for cursor.Next(context.Background()) {
		var transaction *types.Transaction

		err = cursor.Decode(&transaction)
		if err != nil {
			return nil, errors.Wrap(err, "decode transaction")
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (ts *TransactionService) FindUnsent(now int64) ([]*types.Transaction, error) {
	cursor, err := ts.client.DB.Collection(Collection).Find(ts.ctx,
		bson.D{
			{Key: "tx_hash", Value: ""},
			{Key: "from_time", Value: bson.D{{Key: "$lte", Value: now}}},
			{Key: "to_time", Value: bson.D{{Key: "$gte", Value: now}}},
		})
	if err != nil {
		return nil, errors.Wrap(err, "get transactions")
	}

	var transactions []*types.Transaction
	for cursor.Next(context.Background()) {
		var transaction *types.Transaction

		err = cursor.Decode(&transaction)
		if err != nil {
			return nil, errors.Wrap(err, "decode transaction")
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (ts *TransactionService) UpdateTxsHashGasPriceAndTime(vHashes []string, txHash string, gasPrice, now int64) error {
	_, err := ts.client.DB.Collection(Collection).UpdateMany(ts.ctx,
		bson.D{{Key: "v_hash", Value: bson.D{{Key: "$in", Value: vHashes}}}},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "tx_hash", Value: txHash}}},
			{Key: "$set", Value: bson.D{{Key: "gas_price", Value: gasPrice}}},
			{Key: "$set", Value: bson.D{{Key: "send_time", Value: now}}},
		})

	return errors.Wrap(err, "update txs")
}

func (ts *TransactionService) UpdateTxsBlockNumberGasPriceAndTime(vHashes []string, blockNumber, gasPrice, now int64) error {
	_, err := ts.client.DB.Collection(Collection).UpdateMany(ts.ctx,
		bson.D{{Key: "v_hash", Value: bson.D{{Key: "$in", Value: vHashes}}}},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "block_number", Value: blockNumber}}},
			{Key: "$set", Value: bson.D{{Key: "spent_gas", Value: gasPrice}}},
			{Key: "$set", Value: bson.D{{Key: "mined_time", Value: now}}},
		})

	return errors.Wrap(err, "update txs")
}
