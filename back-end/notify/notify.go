package notify

import (
	"light-challenge/config"
	"light-challenge/models"
	"light-challenge/repo"
	"log"
)

type Notifier interface {
	Notify(content interface{}) error
}

type InvoiceNotifier struct {
	repo repo.Repository
}

func (i *InvoiceNotifier) Notify(content interface{}) error {
	invoice := content.(models.Invoice)
	policyCfg := config.Get()

	if invoice.Amount > policyCfg.MaxAmount {
		if invoice.Department == "marketing" {
			log.Println("Sending approval to CMO via mail")
			return nil
		}

		log.Println("Sending approval to CFO via slack")
		return nil
	}

	if invoice.Amount > policyCfg.MeanAmount {
		if invoice.ManagerApproval {
			log.Println("Sending approval request to the Finance manager via mail")
			return nil
		}
	}

	log.Println("Sending approval request to the Finance team via Slack")
	return nil
}

func NewInvoiceNotifier(repo repo.Repository) Notifier {
	return &InvoiceNotifier{
		repo: repo,
	}
}
