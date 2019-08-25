package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/mongo"
	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/service"
	"github.com/miljantekic/ethberlinzwei-maglev/trebuchet/types"
	"github.com/rs/cors"
)

type App struct {
	TransactionService types.TransactionService

	Router *httprouter.Router
}

func MustConfigureApp() *App {
	client := mongo.MustConfigure()

	app := &App{
		TransactionService: mongo.NewTransactionService(client),
	}

	app.Router = NewRouter(app)

	return app
}

func main() {
	app := MustConfigureApp()

	err := app.Start()
	if err != nil {
		panic(err)
	}
}

func (app *App) Start() error {
	handler := cors.AllowAll().Handler(app.Router)

	errCh := make(chan error)
	go func() {
		log.Printf("Listening on port %d...", 80)
		errCh <- http.ListenAndServe(fmt.Sprintf(":%d", 80), handler)
	}()

	go func() {
		batcher := service.NewBatcher(app.TransactionService)
		errCh <- batcher.Start(context.Background())
	}()

	go func() {
		inserter := service.NewInserter(app.TransactionService)
		errCh <- inserter.Start(context.Background())
	}()

	return <-errCh
}
