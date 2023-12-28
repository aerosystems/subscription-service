package monobank

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"time"
)

type Webhook struct {
	InvoiceID     string       `json:"invoiceId"`
	Status        string       `json:"status"`
	FailureReason string       `json:"failureReason"`
	Amount        int          `json:"amount"`
	Ccy           int          `json:"ccy"`
	FinalAmount   int          `json:"finalAmount"`
	CreatedDate   time.Time    `json:"createdDate"`
	ModifiedDate  time.Time    `json:"modifiedDate"`
	Reference     string       `json:"reference"`
	CancelList    []CancelItem `json:"cancelList"`
}

type CancelItem struct {
	Status       string    `json:"status"`
	Amount       int       `json:"amount"`
	Ccy          int       `json:"ccy"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
	ApprovalCode string    `json:"approvalCode"`
	Rrn          string    `json:"rrn"`
	ExtRef       string    `json:"extRef"`
}

func (c *Client) CheckWebhookSignature(bodyBytes []byte, xSignBase64 string) error {
	pubKeyBase64, err := c.getPubKey()
	if err != nil {
		return err
	}
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyBase64)
	if err != nil {
		return err
	}
	signatureBytes, err := base64.StdEncoding.DecodeString(xSignBase64)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(pubKeyBytes)
	if block == nil {
		return errors.New("invalid pem")
	}
	genericPubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pubKey, ok := genericPubKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("invalid public key")
	}
	hash := sha256.Sum256(bodyBytes)
	ok = ecdsa.VerifyASN1(pubKey, hash[:], signatureBytes)
	if !ok {
		return errors.New("invalid signature")
	}
	return nil
}

func (c *Client) GetWebhookFromRequest(bodyBytes []byte) (*Webhook, error) {
	var webhook Webhook
	if err := json.Unmarshal(bodyBytes, &webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}
