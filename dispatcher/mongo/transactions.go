package mongo

import (
	"context"

	"github.com/miljantekic/ethberlinzwei-maglev/dispatcher/types"
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
