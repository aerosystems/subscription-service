package services

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/aerosystems/subs-service/pkg/monobank"
	"github.com/google/uuid"
)

const (
	defaultCcy  = 840 // USD
	redirectUrl = "https://verifire.app/payment/success"
	webHookUrl  = "https://gw.verifire.app/subs/v1/webhook/monobank"
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

func (ps *PaymentServiceImpl) GetPaymentUrl(userUuid uuid.UUID, subscription models.KindSubscription, duration models.DurationSubscription) (string, error) {
	amount, err := ps.priceRepo.GetPrice(subscription, duration)
	if err != nil {
		return "", err
	}
	monoInvoice := &monobank.Invoice{
		Amount:      amount,
		Ccy:         defaultCcy,
		RedirectURL: redirectUrl,
		WebHookURL:  webHookUrl,
	}
	result, err := ps.monobankClient.CreateInvoice(monoInvoice)
	if err != nil {
		return "", err
	}
	invoice := &models.Invoice{
		UserUuid:           userUuid,
		Amount:             amount,
		PaymentMethod:      ps.paymentMethod,
		PaymentStatus:      models.PaymentStatusCreated,
		AcquiringInvoiceId: result.InvoiceId,
	}
	if err := ps.invoiceRepo.Create(invoice); err != nil {
		return "", err
	}
	return result.PageUrl, nil
}
