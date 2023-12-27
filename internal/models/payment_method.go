package models

type PaymentMethod string

const (
	MonobankPaymentMethod PaymentMethod = "monobank"
)

func (p PaymentMethod) String() string {
	return string(p)
}
