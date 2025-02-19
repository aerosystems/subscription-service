package usecases

import (
	"context"
	"fmt"
	"github.com/aerosystems/subscription-service/internal/entities"
	"github.com/google/uuid"
	"time"
)

const (
	titlePage = "Verifire - protect your business"
)

type PaymentUsecase struct {
	invoiceRepo       InvoiceRepository
	priceRepo         PriceRepository
	acquiringRegistry map[entities.PaymentMethod]AcquiringOperations
}

func NewPaymentUsecase(invoiceRepo InvoiceRepository, priceRepo PriceRepository, operations ...AcquiringOperations) *PaymentUsecase {
	if invoiceRepo == nil {
		panic("invoiceRepo is required")
	}
	if priceRepo == nil {
		panic("priceRepo is required")
	}

	registry := map[entities.PaymentMethod]AcquiringOperations{
		entities.UnknownPaymentMethod: &UnknownAcquiring{},
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
	GetPaymentMethod() entities.PaymentMethod
	CreateInvoice(amount int, invoiceUuid, title, description string) (Invoice, error)
	GetWebhookFromRequest(bodyBytes []byte, headers map[string][]string) (Webhook, error)
}

func (ps PaymentUsecase) GetPaymentUrl(ctx context.Context, userUUID uuid.UUID, paymentMethod entities.PaymentMethod, subscription entities.SubscriptionType, duration entities.SubscriptionDuration) (string, error) {
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
	if err = ps.invoiceRepo.Create(ctx, &entities.Invoice{
		CustomerUuid:       userUUID,
		Amount:             amount,
		InvoiceUuid:        invoiceUuid,
		PaymentMethod:      acquiring.GetPaymentMethod(),
		PaymentStatus:      entities.PaymentStatusCreated,
		AcquiringInvoiceId: invoice.AcquiringInvoiceId,
	}); err != nil {
		return "", err
	}
	return invoice.AcquiringPageUrl, nil
}

func (ps PaymentUsecase) ProcessingWebhookPayment(ctx context.Context, paymentMethod entities.PaymentMethod, bodyBytes []byte, headers map[string][]string) error {
	acquiring, ok := ps.acquiringRegistry[paymentMethod]
	if !ok {
		fmt.Errorf("unknown payment method %s", paymentMethod)
	}
	webhook, err := acquiring.GetWebhookFromRequest(bodyBytes, headers)
	if err != nil {
		return err
	}
	invoice, err := ps.invoiceRepo.GetByAcquiringInvoiceId(ctx, webhook.AcquiringInvoiceId)
	if err != nil {
		return err
	}
	if invoice.UpdatedAt.After(webhook.ModifiedDate) { // to prevent not actual webhook
		return nil
	}
	invoice.PaymentStatus = entities.NewPaymentStatus(webhook.Status)
	invoice.UpdatedAt = webhook.ModifiedDate
	if err := ps.invoiceRepo.Update(ctx, invoice); err != nil {
		return err
	}
	return nil
}

func (ps PaymentUsecase) GetPrices(ctx context.Context) map[entities.SubscriptionType]map[entities.SubscriptionDuration]int {
	return ps.priceRepo.GetAll()
}
