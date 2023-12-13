package monobank

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	invoiceCreateUrl = "https://api.monobank.ua/api/merchant/invoice/create"
)

type Invoice struct {
	Amount           int              `json:"amount"`
	Ccy              Ccy              `json:"ccy"`
	MerchantPaymInfo MerchantPaymInfo `json:"merchantPaymInfo"`
	RedirectURL      string           `json:"redirectUrl"`
	WebHookURL       string           `json:"webHookUrl"`
	Validity         int              `json:"validity"`
	PaymentType      string           `json:"paymentType"`
	QrId             string           `json:"qrId"`
	Code             string           `json:"code"`
	SaveCardData     SaveCardData     `json:"saveCardData"`
}

type MerchantPaymInfo struct {
	Reference      string        `json:"reference"`
	Destination    string        `json:"destination"`
	Comment        string        `json:"comment"`
	CustomerEmails []interface{} `json:"customerEmails"`
	BasketOrder    []BasketOrder `json:"basketOrder"`
}

type BasketOrder struct {
	Name      string        `json:"name"`
	Qty       int           `json:"qty"`
	Sum       int           `json:"sum"`
	Icon      string        `json:"icon"`
	Unit      string        `json:"unit"`
	Code      string        `json:"code"`
	Barcode   string        `json:"barcode"`
	Header    string        `json:"header"`
	Footer    string        `json:"footer"`
	Tax       []interface{} `json:"tax"`
	Uktzed    string        `json:"uktzed"`
	Discounts []Discount    `json:"discounts"`
}

type Discount struct {
	Type  string `json:"type"`
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type SaveCardData struct {
	SaveCard bool   `json:"saveCard"`
	WalletId string `json:"walletId"`
}

type InvoiceData struct {
	InvoiceId string `json:"invoiceId"`
	PageUrl   string `json:"pageUrl"`
}

func (c *Client) CreateInvoice(invoice *Invoice) (*InvoiceData, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(invoice); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, invoiceCreateUrl, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Token", c.xToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response InvoiceData
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil

}
