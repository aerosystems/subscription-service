package usecases

import (
	"encoding/json"
	"github.com/aerosystems/subscription-service/internal/entities"
	"github.com/aerosystems/subscription-service/pkg/monobank"
)

const signHeaderName = "X-Sign"

type MonobankAcquiring struct {
	monobankAcquiring *monobank.Acquiring
	paymentMethod     entities.PaymentMethod
	redirectUrl       string
	webHookUrl        string
	currency          monobank.Ccy
}

func NewMonobankAcquiring(monobankClient *monobank.Acquiring, redirectUrl, webHookUrl string, currency monobank.Ccy) *MonobankAcquiring {
	if monobankClient == nil {
		panic("monobankClient is required")
	}
	return &MonobankAcquiring{
		monobankAcquiring: monobankClient,
		paymentMethod:     entities.MonobankPaymentMethod,
		redirectUrl:       redirectUrl,
		webHookUrl:        webHookUrl,
		currency:          currency,
	}
}

func (ms MonobankAcquiring) GetPaymentMethod() entities.PaymentMethod {
	return ms.paymentMethod
}

func (ms MonobankAcquiring) CreateInvoice(amount int, invoiceUuid, title, description string) (Invoice, error) {
	monoInvoice := &monobank.Invoice{
		Amount:      amount,
		Ccy:         ms.currency,
		RedirectURL: ms.redirectUrl,
		WebHookURL:  ms.webHookUrl,
		MerchantPaymInfo: monobank.MerchantPaymInfo{
			Destination: title,
			Comment:     description,
			Reference:   invoiceUuid,
		},
	}
	invoice, err := ms.monobankAcquiring.CreateInvoice(monoInvoice)
	if err != nil {
		return Invoice{}, err
	}
	return Invoice{invoice.InvoiceId, invoice.PageUrl}, nil
}

func (ms MonobankAcquiring) GetWebhookFromRequest(bodyBytes []byte, headers map[string][]string) (Webhook, error) {
	xSignBase64 := headers[signHeaderName][0]
	if err := ms.monobankAcquiring.CheckWebhookSignature(bodyBytes, xSignBase64); err != nil {
		return Webhook{}, err
	}
	var webhook monobank.Webhook
	if err := json.Unmarshal(bodyBytes, &webhook); err != nil {
		return Webhook{}, err
	}
	return Webhook{
		AcquiringInvoiceId: webhook.InvoiceID,
		Status:             webhook.Status,
		ModifiedDate:       webhook.ModifiedDate,
	}, nil
}

func (ms MonobankAcquiring) ConvertStatus(status string) entities.PaymentStatus {
	switch status {
	case monobank.InvoiceStatusProcessing:
		return entities.PaymentStatusPending
	case monobank.InvoiceStatusSuccess:
		return entities.PaymentStatusPaid
	default:
		return entities.PaymentStatusFailed
	}

}
