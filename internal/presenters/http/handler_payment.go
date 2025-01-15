package HTTPServer

type PaymentHandler struct {
	*BaseHandler
	paymentUsecase PaymentUsecase
}

func NewPaymentHandler(
	baseHandler *BaseHandler,
	paymentUsecase PaymentUsecase,
) *PaymentHandler {
	if baseHandler == nil {
		panic("baseHandler is required")
	}
	if paymentUsecase == nil {
		panic("paymentUsecase is required")
	}
	return &PaymentHandler{
		BaseHandler:    baseHandler,
		paymentUsecase: paymentUsecase,
	}
}
