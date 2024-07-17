package models

type PaymentMethod struct {
	slug string
}

var (
	UnknownPaymentMethod  = PaymentMethod{"unknown"}
	MonobankPaymentMethod = PaymentMethod{"monobank"}
)

func (p PaymentMethod) String() string {
	return p.slug
}

func NewPaymentMethod(method string) PaymentMethod {
	switch method {
	case MonobankPaymentMethod.String():
		return MonobankPaymentMethod
	default:
		return UnknownPaymentMethod
	}
}

type PaymentStatus struct {
	slug string
}

var (
	UnknownPaymentStatus = PaymentStatus{"unknown"}
	PaymentStatusCreated = PaymentStatus{"created"}
	PaymentStatusPending = PaymentStatus{"pending"}
	PaymentStatusPaid    = PaymentStatus{"paid"}
	PaymentStatusFailed  = PaymentStatus{"failed"}
)

func (p PaymentStatus) String() string {
	return p.slug
}

func NewPaymentStatus(status string) PaymentStatus {
	switch status {
	case PaymentStatusCreated.String():
		return PaymentStatusCreated
	case PaymentStatusPending.String():
		return PaymentStatusPending
	case PaymentStatusPaid.String():
		return PaymentStatusPaid
	case PaymentStatusFailed.String():
		return PaymentStatusFailed
	default:
		return UnknownPaymentStatus
	}
}
