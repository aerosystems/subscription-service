package monobank

type Invoice struct {
	Amount           int              `json:"amount"`
	Ccy              int              `json:"ccy"`
	MerchantPaymInfo MerchantPaymInfo `json:"merchantPaymInfo"`
	RedirectURL      string           `json:"redirectUrl"`
	WebHookURL       string           `json:"webHookUrl"`
	Validity         int              `json:"validity"`
	PaymentType      string           `json:"paymentType"`
	QrID             string           `json:"qrId"`
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
	WalletID string `json:"walletId"`
}
