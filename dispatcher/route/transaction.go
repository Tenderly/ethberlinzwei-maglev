package route

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/miljantekic/ethberlinzwei-maglev/dispatcher/route/helper"
	"github.com/miljantekic/ethberlinzwei-maglev/dispatcher/types"
	"github.com/pkg/errors"
)

func Insert(transactionService types.TransactionService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var tx types.Transaction

		err := json.NewDecoder(r.Body).Decode(&tx)
		if err != nil {
			//log.ErrorErr(err, "could not read request")

			helper.BadRequestError(w, errors.New("could not read request"))

			return
		}

		tx.VHash = "0x" + helper.RandStringRunes(32)

		err = transactionService.Insert(&tx)
		if err != nil {
			//log.ErrorErr(err, "could not insert transaction")

			helper.BadRequestError(w, errors.New("could not insert transaction"))

			return
		}

		helper.SendResponse(w, 200, tx)
	}
}

func Find(transactionService types.TransactionService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		txs, err := transactionService.Find()
		if err != nil {
			//log.ErrorErr(err, "could not find transaction")

			helper.BadRequestError(w, errors.New("could not find transaction"))

			return
		}

		helper.SendResponse(w, 200, txs)
	}
}
