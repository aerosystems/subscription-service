package services

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/aerosystems/subs-service/pkg/monobank"
	"github.com/google/uuid"
)

type PaymentServiceImpl struct {
	paymentMethod  models.PaymentMethod
	invoiceRepo    repository.InvoiceRepository
	priceRepo      repository.PriceRepository
	monobankClient *monobank.Client
}

func NewPaymentServiceImpl(invoiceRepo repository.InvoiceRepository, priceRepo repository.PriceRepository, monobankClient *monobank.Client) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		invoiceRepo:    invoiceRepo,
		priceRepo:      priceRepo,
		monobankClient: monobankClient,
	}
}

func (ps *PaymentServiceImpl) SetPaymentMethod(paymentMethod string) error {
	switch paymentMethod {
	case models.MonobankPaymentMethod.String():
		ps.paymentMethod = models.MonobankPaymentMethod
	}
	return nil
}

func (ps *PaymentServiceImpl) CreateInvoice(userUuid uuid.UUID, amount int) (*models.Invoice, error) {
	invoice := &models.Invoice{
		UserUuid:      userUuid,
		Amount:        amount,
		PaymentMethod: ps.paymentMethod,
	}
	monoInvoice := &monobank.Invoice{
		Amount: amount,
	}
	_, err := ps.monobankClient.CreateInvoice(monoInvoice)
	if err != nil {
		return nil, err
	}
	if err := ps.invoiceRepo.Create(invoice); err != nil {
		return nil, err
	}
	return invoice, nil
}

func (ps *PaymentServiceImpl) GetPrice(kindSubscription, durationSubscription string) (int, error) {
	return ps.priceRepo.GetPrice(kindSubscription, durationSubscription)
}
