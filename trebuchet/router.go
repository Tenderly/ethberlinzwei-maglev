package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/route"
)

func NewRouter(app *App) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {})
	router.POST("/insert", route.Insert(app.TransactionService))
	router.GET("/find", route.Find(app.TransactionService))

	return router
}
