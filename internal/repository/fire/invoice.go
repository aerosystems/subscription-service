package fire

import (
	"cloud.google.com/go/firestore"
	"context"
)

type InvoiceRepo struct {
	client *firestore.Client
}

func NewInvoiceRepo(client *firestore.Client) *InvoiceRepo {
	return &InvoiceRepo{
		client: client,
	}
}

func (r *InvoiceRepo) GetByUserUuid(ctx context.Context, userUuid string) ([]map[string]interface{}, error) {
	var invoices []map[string]interface{}

	iter := r.client.Collection("invoices").Where("UserUuid", "==", userUuid).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		invoices = append(invoices, doc.Data())
	}

	return invoices, nil
}

func (r *InvoiceRepo) GetByUuid(ctx context.Context, uuid string) (map[string]interface{}, error) {
	docRef := r.client.Collection("invoices").Doc(uuid)
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	return doc.Data(), nil
}

func (r *InvoiceRepo) Create(ctx context.Context, invoice map[string]interface{}) error {
	_, err := r.client.Collection("invoices").Doc(invoice["Uuid"].(string)).Set(ctx, invoice)
	if err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepo) Update(ctx context.Context, invoice map[string]interface{}) error {
	_, err := r.client.Collection("invoices").Doc(invoice["Uuid"].(string)).Set(ctx, invoice)
	if err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepo) Delete(ctx context.Context, invoice map[string]interface{}) error {
	_, err := r.client.Collection("invoices").Doc(invoice["Uuid"].(string)).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
