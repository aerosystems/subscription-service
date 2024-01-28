package payment

import (
	"fmt"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/internal/repository/pg"
	"github.com/google/uuid"
	"time"
)

const (
	titlePage = "Verifire - protect your business"
)

type ServiceImpl struct {
	acquiring   AcquiringOperations
	invoiceRepo pg.InvoiceRepository
	priceRepo   pg.PriceRepository
	strategies  map[models.PaymentMethod]AcquiringOperations
}

func NewServiceImpl(invoiceRepo pg.InvoiceRepository, priceRepo pg.PriceRepository, strategies map[models.PaymentMethod]AcquiringOperations) *ServiceImpl {
	return &ServiceImpl{
		invoiceRepo: invoiceRepo,
		priceRepo:   priceRepo,
		strategies:  strategies,
	}
}

type Invoice struct {
	AcquiringInvoiceId string
	AcquiringPageUrl   string
}

type Webhook struct {
	AcquiringInvoiceId string
	Status             string
	ModifiedDate       time.Time
}

type AcquiringOperations interface {
	GetPaymentMethod() models.PaymentMethod
	CreateInvoice(amount int, invoiceUuid, title, description string) (Invoice, error)
	GetWebhookFromRequest(bodyBytes []byte, headers map[string][]string) (Webhook, error)
}

func (ps *ServiceImpl) SetPaymentMethod(paymentMethod models.PaymentMethod) error {
	if _, ok := ps.strategies[paymentMethod]; !ok {
		return fmt.Errorf("invalid payment method")
	}
	ps.acquiring = ps.strategies[paymentMethod]
	return nil
}

func (ps *ServiceImpl) GetPaymentUrl(userUuid uuid.UUID, subscription models.KindSubscription, duration models.DurationSubscription) (string, error) {
	amount, err := ps.priceRepo.GetPrice(subscription, duration)
	if err != nil {
		return "", err
	}
	invoiceUuid := uuid.New()
	description := fmt.Sprintf("Subscription %s with %s payment", subscription.String(), duration.String())
	invoice, err := ps.acquiring.CreateInvoice(amount, invoiceUuid.String(), titlePage, description)
	if err != nil {
		return "", err
	}
	if err := ps.invoiceRepo.Create(&models.Invoice{
		UserUuid:           userUuid,
		Amount:             amount,
		InvoiceUuid:        invoiceUuid,
		PaymentMethod:      ps.acquiring.GetPaymentMethod(),
		PaymentStatus:      models.PaymentStatusCreated,
		AcquiringInvoiceId: invoice.AcquiringInvoiceId,
	}); err != nil {
		return "", err
	}
	return invoice.AcquiringPageUrl, nil
}

func (ps *ServiceImpl) ProcessingWebhookPayment(bodyBytes []byte, headers map[string][]string) error {
	webhook, err := ps.acquiring.GetWebhookFromRequest(bodyBytes, headers)
	if err != nil {
		return err
	}
	invoice, err := ps.invoiceRepo.GetByAcquiringInvoiceId(webhook.AcquiringInvoiceId)
	if err != nil {
		return err
	}
	if invoice.UpdatedAt.After(webhook.ModifiedDate) { // to prevent not actual webhook
		return nil
	}
	invoice.PaymentStatus = models.PaymentStatus(webhook.Status)
	invoice.UpdatedAt = webhook.ModifiedDate
	if err := ps.invoiceRepo.Update(invoice); err != nil {
		return err
	}
	return nil
}
