package services

import "github.com/aerosystems/subs-service/internal/models"

type MonobankPaymentService struct {
	invoiceRepo models.InvoiceRepository
}

func NewMonobankPaymentService(invoiceRepo models.InvoiceRepository) *MonobankPaymentService {
	return &MonobankPaymentService{
		invoiceRepo: invoiceRepo,
	}
}

func (m *MonobankPaymentService) CreateInvoice(invoice *models.Invoice) error {
	return m.invoiceRepo.Create(invoice)
}
