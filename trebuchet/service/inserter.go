package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/route/helper"
	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/types"
	"go.mongodb.org/mongo-driver/bson"
)

var addresses = []string{
	"0x0dfcb229eab0b787beb035314bc51550dbdbc957",
	"0x07a7ced65a227d6e6f762a4cbefb68dfb4e9b1ce",
	"0xa0c3c24298a4c555d94132c1652e3f5ab24140c6",
	"0xf23d5e7d9ca123dadb80edb9240ee1de6ce1f8ec",
	"0x121ef61dbbdde01ccc2323055328b26dc5e85739",
}

type Inserter struct {
	TransactionService types.TransactionService
}

func NewInserter(transactionService types.TransactionService) *Inserter {
	return &Inserter{
		TransactionService: transactionService,
	}
}

func (b Inserter) Start(ctx context.Context) error {
	log.Println("Started inserter")
	for {

		addr1 := addresses[rand.Intn(5)]
		addr2 := addr1

		for addr1 == addr2 {
			addr2 = addresses[rand.Intn(5)]
		}

		err := b.TransactionService.InsertRaw(bson.D{
			bson.E{Key: "v_hash", Value: "0x" + helper.RandStringRunes(64)},
			bson.E{Key: "v_nonce", Value: 0},

			bson.E{Key: "from_time", Value: 0},
			bson.E{Key: "to_time", Value: time.Now().Unix() + 300},

			bson.E{Key: "operation_type", Value: 0},
			bson.E{Key: "identity", Value: addr1},

			bson.E{Key: "to", Value: addr2},
			bson.E{Key: "value", Value: 1},
			bson.E{Key: "payload", Value: "0x"},

			bson.E{Key: "tx_hash", Value: ""},
			bson.E{Key: "gas", Value: 100000},

			bson.E{Key: "estimateGas", Value: 31544},
		})

		if err != nil {
			log.Println(fmt.Sprintf("error inserting tx, %s", err))
		}

		time.Sleep(time.Duration(time.Second.Nanoseconds() + rand.Int63n(time.Second.Nanoseconds())))
	}
}
