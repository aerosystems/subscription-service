package usecases

import (
	"context"
	"fmt"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/google/uuid"
	"time"
)

const (
	titlePage = "Verifire - protect your business"
)

type PaymentUsecase struct {
	invoiceRepo       InvoiceRepository
	priceRepo         PriceRepository
	acquiringRegistry map[models.PaymentMethod]AcquiringOperations
}

func NewPaymentUsecase(invoiceRepo InvoiceRepository, priceRepo PriceRepository, operations ...AcquiringOperations) *PaymentUsecase {
	if invoiceRepo == nil {
		panic("invoiceRepo is required")
	}
	if priceRepo == nil {
		panic("priceRepo is required")
	}

	registry := map[models.PaymentMethod]AcquiringOperations{
		models.UnknownPaymentMethod: &UnknownAcquiring{},
	}

	for _, op := range operations {
		registry[op.GetPaymentMethod()] = op
	}

	return &PaymentUsecase{
		invoiceRepo:       invoiceRepo,
		priceRepo:         priceRepo,
		acquiringRegistry: registry,
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

func (ps PaymentUsecase) GetPaymentUrl(userUUID uuid.UUID, paymentMethod models.PaymentMethod, subscription models.SubscriptionType, duration models.SubscriptionDuration) (string, error) {
	amount, err := ps.priceRepo.GetPrice(subscription, duration)
	if err != nil {
		return "", err
	}
	invoiceUuid := uuid.New()
	description := fmt.Sprintf("Subscription %s with %s payment", subscription.String(), duration.String())
	acquiring, ok := ps.acquiringRegistry[paymentMethod]
	if !ok {
		return "", fmt.Errorf("unknown payment method %s", paymentMethod)
	}
	invoice, err := acquiring.CreateInvoice(amount, invoiceUuid.String(), titlePage, description)
	if err != nil {
		return "", err
	}
	if err = ps.invoiceRepo.Create(context.Background(), &models.Invoice{
		CustomerUuid:       userUUID,
		Amount:             amount,
		InvoiceUuid:        invoiceUuid,
		PaymentMethod:      acquiring.GetPaymentMethod(),
		PaymentStatus:      models.PaymentStatusCreated,
		AcquiringInvoiceId: invoice.AcquiringInvoiceId,
	}); err != nil {
		return "", err
	}
	return invoice.AcquiringPageUrl, nil
}

func (ps PaymentUsecase) ProcessingWebhookPayment(paymentMethod models.PaymentMethod, bodyBytes []byte, headers map[string][]string) error {
	acquiring, ok := ps.acquiringRegistry[paymentMethod]
	if !ok {
		fmt.Errorf("unknown payment method %s", paymentMethod)
	}
	webhook, err := acquiring.GetWebhookFromRequest(bodyBytes, headers)
	if err != nil {
		return err
	}
	invoice, err := ps.invoiceRepo.GetByAcquiringInvoiceId(context.Background(), webhook.AcquiringInvoiceId)
	if err != nil {
		return err
	}
	if invoice.UpdatedAt.After(webhook.ModifiedDate) { // to prevent not actual webhook
		return nil
	}
	invoice.PaymentStatus = models.NewPaymentStatus(webhook.Status)
	invoice.UpdatedAt = webhook.ModifiedDate
	if err := ps.invoiceRepo.Update(context.Background(), invoice); err != nil {
		return err
	}
	return nil
}

func (ps PaymentUsecase) GetPrices() map[models.SubscriptionType]map[models.SubscriptionDuration]int {
	return ps.priceRepo.GetAll()
}
