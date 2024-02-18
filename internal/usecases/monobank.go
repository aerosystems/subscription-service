package usecases

import (
	"encoding/json"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/aerosystems/subs-service/pkg/monobank"
)

type MonobankStrategy struct {
	monobankAcquiring *monobank.Acquiring
	paymentMethod     models.PaymentMethod
	redirectUrl       string
	webHookUrl        string
	currency          monobank.Ccy
}

func NewMonobankStrategy(monobankClient *monobank.Acquiring, redirectUrl, webHookUrl string, currency monobank.Ccy) *MonobankStrategy {
	return &MonobankStrategy{
		monobankAcquiring: monobankClient,
		paymentMethod:     models.MonobankPaymentMethod,
		redirectUrl:       redirectUrl,
		webHookUrl:        webHookUrl,
		currency:          currency,
	}
}

func (ms MonobankStrategy) GetPaymentMethod() models.PaymentMethod {
	return ms.paymentMethod
}

func (ms MonobankStrategy) CreateInvoice(amount int, invoiceUuid, title, description string) (Invoice, error) {
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

func (ms MonobankStrategy) GetWebhookFromRequest(bodyBytes []byte, headers map[string][]string) (Webhook, error) {
	xSignBase64 := headers["X-Sign"][0]
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

func (ms MonobankStrategy) ConvertStatus(status string) models.PaymentStatus {
	switch status {
	case monobank.InvoiceStatusProcessing:
		return models.PaymentStatusPending
	case monobank.InvoiceStatusSuccess:
		return models.PaymentStatusPaid
	default:
		return models.PaymentStatusFailed
	}

}
