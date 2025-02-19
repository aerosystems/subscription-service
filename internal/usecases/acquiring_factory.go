package usecases

import (
	"errors"
	"github.com/aerosystems/subscription-service/internal/entities"
)

type UnknownAcquiring struct{}

func (us UnknownAcquiring) GetPaymentMethod() entities.PaymentMethod {
	return entities.UnknownPaymentMethod
}

func (us UnknownAcquiring) CreateInvoice(amount int, invoiceUuid, title, description string) (Invoice, error) {
	return Invoice{}, errors.New("unknown payment method")
}

func (us UnknownAcquiring) GetWebhookFromRequest(bodyBytes []byte, headers map[string][]string) (Webhook, error) {
	return Webhook{}, errors.New("unknown payment method")
}
