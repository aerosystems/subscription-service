package usecases

import (
	"errors"
	"github.com/aerosystems/subscription-service/internal/models"
)

type UnknownStrategy struct{}

func NewUnknownStrategy() *UnknownStrategy {
	return &UnknownStrategy{}
}

func (us UnknownStrategy) GetPaymentMethod() models.PaymentMethod {
	return models.UnknownPaymentMethod
}

func (us UnknownStrategy) CreateInvoice(amount int, invoiceUuid, title, description string) (Invoice, error) {
	return Invoice{}, errors.New("unknown payment method")
}

func (us UnknownStrategy) GetWebhookFromRequest(bodyBytes []byte, headers map[string][]string) (Webhook, error) {
	return Webhook{}, errors.New("unknown payment method")
}
