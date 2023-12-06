package services

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/pkg/monobank"
)

type PaymentService interface {
	SetPaymentMethod(paymentMethod string) error
	CreateInvoice(invoice *models.Invoice) error
}

type PaymentServiceImpl struct {
	paymentMethod  models.PaymentMethod
	invoiceRepo    models.InvoiceRepository
	monobankClient *monobank.Client
}

func NewPaymentServiceImpl(invoiceRepo models.InvoiceRepository, monobankClient *monobank.Client) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		invoiceRepo:    invoiceRepo,
		monobankClient: monobankClient,
	}
}

func (ps *PaymentServiceImpl) SetPaymentMethod(paymentMethod string) error {
	switch paymentMethod {
	case "monobank":
		ps.paymentMethod = models.MonobankPaymentMethod
	}
	return nil
}

func (ps *PaymentServiceImpl) CreateInvoice(invoice *models.Invoice) error {
	return ps.monobankClient.CreateInvoice(nil)
}
