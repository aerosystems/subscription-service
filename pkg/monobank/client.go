package monobank

type MonoClient interface {
	CreateInvoice(invoice *Invoice) (*InvoiceData, error)
	CheckWebhookSignature(bodyBytes []byte, xSignBase64 string) error
	GetWebhookFromRequest(bodyBytes []byte) (*Webhook, error)
}

type Client struct {
	xToken string
}

func NewClient(xToken string) *Client {
	return &Client{
		xToken: xToken,
	}
}
