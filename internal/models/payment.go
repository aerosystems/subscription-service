package models

type PaymentMethod string

const (
	MonobankPaymentMethod PaymentMethod = "monobank"
)

func (p PaymentMethod) String() string {
	return string(p)
}

type PaymentStatus string

const (
	PaymentStatusCreated PaymentStatus = "created"
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusFailed  PaymentStatus = "failed"
)

func (p PaymentStatus) String() string {
	return string(p)
}
