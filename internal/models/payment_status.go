package models

type PaymentStatus string

const (
	PaymentStatusCreated PaymentStatus = "created"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusFailed  PaymentStatus = "failed"
)

func (p PaymentStatus) String() string {
	return string(p)
}
