package HTTPServer

type SubscriptionHandler struct {
	*BaseHandler
	subscriptionUsecase SubscriptionUsecase
}

func NewSubscriptionHandler(baseHandler *BaseHandler, subscriptionUsecase SubscriptionUsecase) *SubscriptionHandler {
	return &SubscriptionHandler{baseHandler, subscriptionUsecase}
}
