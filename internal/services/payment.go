package services

import (
	"fmt"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository"
	"github.com/aerosystems/subs-service/pkg/monobank"
	"github.com/google/uuid"
)

const (
	redirectUrl = "https://verifire.app/payment/success"
	webHookUrl  = "https://gw.verifire.app/subs/v1/webhook/monobank"
)

type PaymentServiceImpl struct {
	paymentMethod  models.PaymentMethod
	invoiceRepo    repository.InvoiceRepository
	priceRepo      repository.PriceRepository
	monobankClient monobank.MonoClient
}

func NewPaymentServiceImpl(invoiceRepo repository.InvoiceRepository, priceRepo repository.PriceRepository, monobankClient monobank.MonoClient) *PaymentServiceImpl {
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
	invoiceUuid := uuid.New()
	monoInvoice := &monobank.Invoice{
		Amount:      amount,
		Ccy:         monobank.USD,
		RedirectURL: redirectUrl,
		WebHookURL:  webHookUrl,
		MerchantPaymInfo: monobank.MerchantPaymInfo{
			Destination: "Verifire - protect your business",
			Comment: fmt.Sprintf("Subscription %s with %s payment",
				subscription.String(),
				duration.String(),
			),
			Reference: invoiceUuid.String(),
		},
	}
	result, err := ps.monobankClient.CreateInvoice(monoInvoice)
	if err != nil {
		return "", err
	}
	invoice := &models.Invoice{
		UserUuid:           userUuid,
		Amount:             amount,
		InvoiceUuid:        invoiceUuid,
		PaymentMethod:      ps.paymentMethod,
		PaymentStatus:      models.PaymentStatusCreated,
		AcquiringInvoiceId: result.InvoiceId,
	}
	if err := ps.invoiceRepo.Create(invoice); err != nil {
		return "", err
	}
	return result.PageUrl, nil
}

func (ps *PaymentServiceImpl) ProcessingWebhookPayment(userUuid uuid.UUID, invoiceUuid uuid.UUID, acquiringInvoiceId string) error {
	return nil
}
