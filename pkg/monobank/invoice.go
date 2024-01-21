package monobank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

const (
	invoiceCreateUrl = "https://api.monobank.ua/api/merchant/invoice/create"
)

type ErrorResponse struct {
	ErrCode string `json:"errCode"`
	ErrText string `json:"errText"`
}

type Invoice struct {
	Amount           int              `json:"amount"`
	Ccy              Ccy              `json:"ccy"`
	MerchantPaymInfo MerchantPaymInfo `json:"merchantPaymInfo"`
	RedirectURL      string           `json:"redirectUrl"`
	WebHookURL       string           `json:"webHookUrl"`
	Validity         int              `json:"validity,omitempty"`
	PaymentType      string           `json:"paymentType,omitempty"`
	QrId             string           `json:"qrId,omitempty"`
	Code             string           `json:"code,omitempty"`
	SaveCardData     SaveCardData     `json:"saveCardData,omitempty"`
}

type MerchantPaymInfo struct {
	Reference      string        `json:"reference"`
	Destination    string        `json:"destination"`
	Comment        string        `json:"comment"`
	CustomerEmails []interface{} `json:"customerEmails,omitempty"`
	BasketOrder    []BasketOrder `json:"basketOrder,omitempty"`
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

func (a Acquiring) CreateInvoice(invoice *Invoice) (*InvoiceData, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(invoice); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, invoiceCreateUrl, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Token", a.xToken)
	req.Header.Set("Content-Type", "application/json")
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
		return nil, err
	}

	fmt.Println(string(dump))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("monobank request error code: %s, error text: %s", errorResponse.ErrCode, errorResponse.ErrText)
	}
	defer resp.Body.Close()
	var response InvoiceData
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
