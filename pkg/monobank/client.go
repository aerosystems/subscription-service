package monobank

type SDK interface {
	CreateInvoice(request *Invoice) error
}

type Client struct {
	xToken string
}

func NewClient(xToken string) *Client {
	return &Client{
		xToken: xToken,
	}
}
