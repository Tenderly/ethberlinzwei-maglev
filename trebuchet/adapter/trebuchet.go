package adapter

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/types"
	"github.com/pkg/errors"
)

const address = "0x5972ad94694b275a5916f357fc97755bd9de7c7a"

type Executor struct {
	contract *Trebuchet
}

func MustConfigureExecutor(client *ethclient.Client) *Executor {
	thebuchet, err := NewTrebuchet(common.HexToAddress(address), client)
	if err != nil {
		panic(errors.Wrap(err, "generate thebuchet adapter"))
	}

	return &Executor{
		contract: thebuchet,
	}
}

func (e Executor) Send(txs []*types.Transaction) (common.Hash, int64, error) {
	pk, _ := crypto.HexToECDSA("FD5F9928F7EFC0BB2A2534B9BEBDD5D94AD261A920BF1161242B899FCE42F8EA")

	var vHashes [][32]byte

	var data [][]byte
	for _, tx := range txs {
		var datum []byte

		tmpVHash := hexutil.MustDecode(tx.VHash)
		var vHash [32]byte
		copy(vHash[:], tmpVHash)

		vHashes = append(vHashes, vHash)
		datum = append(datum, byte(255), byte(255), byte(255))

		datum = append(datum, common.HexToAddress(tx.Identity).Bytes()...)

		datum = append(datum, []byte{9, 197, 234, 190}...)
		datum = append(datum, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32}...)
		datum = append(datum, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 33}...)

		if tx.OperationType == 0 {
			datum = append(datum, byte(0))
		} else {
			datum = append(datum, byte(1))
		}

		datum = append(datum, common.HexToAddress(tx.To).Bytes()...)
		datum = append(datum, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}...)

		data = append(data, datum)
	}

	txOpts := *bind.NewKeyedTransactor(pk)
	txOpts.GasLimit = 1000000
	txOpts.GasPrice = big.NewInt(1000000000)

	session := &TrebuchetSession{
		Contract:     e.contract,
		CallOpts:     bind.CallOpts{},
		TransactOpts: txOpts,
	}

	ptx, err := session.Execute(data)
	if err != nil {
		return common.Hash{}, 0, errors.Wrap(err, "execute tx")
	}

	for _, tx := range txs {
		tx.TxHash = ptx.Hash().String()
		tx.GasPrice = 1000000000
	}

	return ptx.Hash(), 1000000000, nil
}
