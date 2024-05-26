package repo

import (
	"light-challenge/models"
	"log"

	"github.com/go-pg/pg"
)

type Repository interface {
	Insert(model interface{}) error
}

type Invoice struct {
	db *pg.DB
}

func (i *Invoice) Insert(invoice interface{}) error {
	invoiceModel := invoice.(models.Invoice)
	err := i.db.Insert(&invoiceModel)

	if err != nil {
		log.Printf("Failed to insert an invoice due to: %s\n", err.Error())
		return err
	}

	log.Printf("Invoice added\n")
	return nil
}

func NewInvoiceRepo(db *pg.DB) Repository {
	return &Invoice{
		db: db,
	}
}
