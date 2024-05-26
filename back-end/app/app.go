package app

import (
	"fmt"
	"light-challenge/controller"
	"light-challenge/db"
	"light-challenge/repo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type app struct {
}

func (a *app) Start() error {
	dbConn, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	invoiceRepo := repo.NewInvoiceRepo(dbConn)
	invoiceController := controller.NewInvoice(invoiceRepo)
	router := mux.NewRouter()

	router.HandleFunc("/invoice", invoiceController.Handle).Methods("POST")

	port := 8080
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("localhost:%v", port),
	}

	log.Printf("Starting app on port %v\n", port)

	return server.ListenAndServe()

}

func New() app {
	return app{}
}
