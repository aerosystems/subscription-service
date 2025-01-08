package HTTPServer

type PaymentHandler struct {
	*BaseHandler
	paymentUsecase PaymentUsecase
}

func NewPaymentHandler(
	baseHandler *BaseHandler,
	paymentUsecase PaymentUsecase,
) *PaymentHandler {
	return &PaymentHandler{
		BaseHandler:    baseHandler,
		paymentUsecase: paymentUsecase,
	}
}
