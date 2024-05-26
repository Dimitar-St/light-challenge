package controller

import (
	"encoding/json"
	"light-challenge/models"
	"light-challenge/notify"
	"light-challenge/repo"
	"log"
	"net/http"
)

type Invoice struct {
	repo     repo.Repository
	notifier notify.InvoiceNotifier
}

func (i *Invoice) Handle(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	decoder := json.NewDecoder(request.Body)

	var invoice models.Invoice = models.Invoice{}

	err := decoder.Decode(&invoice)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte("Failed to parse invoice"))
		log.Printf("invalid invoice %v\n", err)
		return
	}

	log.Println("Handling request")

	err = i.repo.Insert(invoice)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte("Failed to save invoice"))
		return
	}

	i.notifier.Notify(invoice)
}

func NewInvoice(repo repo.Repository) Invoice {
	return Invoice{
		repo: repo,
	}
}
