package services

import "github.com/aerosystems/subs-service/internal/models"

type InvoiceService interface {
	CreateInvoice(invoice *models.Invoice) error
}
