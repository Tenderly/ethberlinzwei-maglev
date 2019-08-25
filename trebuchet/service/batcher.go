package service

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/adapter"
	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/types"
	"github.com/pkg/errors"
)

const httpProvider = "https://rinkeby.tenderly.dev"

type Batcher struct {
	TransactionService types.TransactionService
	Executor           *adapter.Executor
	Client             *ethclient.Client
}

func NewBatcher(transactionService types.TransactionService) *Batcher {
	client, err := ethclient.Dial(httpProvider)
	if err != nil {
		panic(errors.Wrap(err, "dial client"))
	}

	return &Batcher{
		TransactionService: transactionService,
		Executor:           adapter.MustConfigureExecutor(client),
		Client:             client,
	}
}

func (b Batcher) Start(ctx context.Context) error {
	log.Println("Started batcher")
	for {
		txs, err := b.TransactionService.FindUnsent(time.Now().Unix())
		if err != nil {
			log.Println("error fetching transactions to send")
			time.Sleep(30 * time.Second)
			continue
		}

		if len(txs) <= 20 {
			time.Sleep(5 * time.Second)
			continue
		}

		hash, gasPrice, err := b.Executor.Send(txs)
		if err != nil {
			log.Println("error sending transaction")
			continue
		}

		var vHashes []string
		for _, tx := range txs {
			vHashes = append(vHashes, tx.VHash)
		}

		err = b.TransactionService.UpdateTxsHashGasPriceAndTime(vHashes, hash.String(), gasPrice, time.Now().UnixNano())
		if err != nil {
			log.Println("error updating tx hashes and gas price for txs vhash")
			continue
		}

		go b.wait(hash, vHashes)
	}
}

func (b Batcher) wait(hash common.Hash, vHashes []string) {
	receipt, err := b.Client.TransactionReceipt(context.Background(), hash)

	for err != nil {
		time.Sleep(500 * time.Millisecond)
		receipt, err = b.Client.TransactionReceipt(context.Background(), hash)
	}

	err = b.TransactionService.UpdateTxsBlockNumberGasPriceAndTime(vHashes, receipt.BlockNumber.Int64(), int64(receipt.GasUsed)/int64(len(vHashes)), time.Now().UnixNano())
	if err != nil {
		log.Println("error updating block number for txs vhash")
	}

	return
}
