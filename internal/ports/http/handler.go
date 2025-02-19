package HTTPServer

type Handler struct {
	subscriptionUsecase SubscriptionUsecase
	paymentUsecase      PaymentUsecase
}

func NewHandler(subscriptionUsecase SubscriptionUsecase, paymentUsecase PaymentUsecase) *Handler {
	return &Handler{
		subscriptionUsecase: subscriptionUsecase,
		paymentUsecase:      paymentUsecase,
	}
}
