package fire

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/aerosystems/subscription-service/internal/models"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"time"
)

const (
	collectionName = "invoices"
)

type InvoiceRepo struct {
	client *firestore.Client
}

func NewInvoiceRepo(client *firestore.Client) *InvoiceRepo {
	return &InvoiceRepo{
		client: client,
	}
}

type InvoiceFire struct {
	Amount             int       `firestore:"amount"`
	CustomerUuid       string    `firestore:"customer_uuid"`
	InvoiceUuid        string    `firestore:"invoice_uuid"`
	PaymentMethod      string    `firestore:"payment_method"`
	AcquiringInvoiceId string    `firestore:"acquiring_invoice_id"`
	PaymentStatus      string    `firestore:"payment_status"`
	CreatedAt          time.Time `firestore:"created_at"`
	UpdatedAt          time.Time `firestore:"updated_at"`
}

func (i *InvoiceFire) ToModel() (models.Invoice, error) {
	customerId, err := uuid.Parse(i.CustomerUuid)
	if err != nil {
		return models.Invoice{}, err
	}
	invoiceUuid, err := uuid.Parse(i.InvoiceUuid)
	if err != nil {
		return models.Invoice{}, err
	}

	return models.Invoice{
		Amount:             i.Amount,
		CustomerUuid:       customerId,
		InvoiceUuid:        invoiceUuid,
		PaymentMethod:      models.NewPaymentMethod(i.PaymentMethod),
		AcquiringInvoiceId: i.AcquiringInvoiceId,
		PaymentStatus:      models.NewPaymentStatus(i.PaymentStatus),
		CreatedAt:          i.CreatedAt,
		UpdatedAt:          i.UpdatedAt,
	}, nil
}

func (r *InvoiceRepo) Create(ctx context.Context, invoice *models.Invoice) error {
	_, _, err := r.client.Collection(collectionName).Add(ctx, invoice)
	if err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepo) GetByAcquiringInvoiceId(ctx context.Context, acquiringInvoiceId string) (*models.Invoice, error) {
	c, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var invoiceFire InvoiceFire
	iter := r.client.Collection(collectionName).Where("acquiring_invoice_id", "==", acquiringInvoiceId).Documents(c)
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		doc.DataTo(&invoiceFire)
	}
	invoice, err := invoiceFire.ToModel()
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *InvoiceRepo) Update(ctx context.Context, invoice *models.Invoice) error {
	c, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	_, err := r.client.Collection(collectionName).Doc(invoice.InvoiceUuid.String()).Set(c, invoice)
	if err != nil {
		return err
	}
	return nil
}
