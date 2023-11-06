package RPCServer

type SubsRPCPayload struct {
	UserId int
	Kind   string
}

func (ss *SubsServer) CreateFreeTrial(payload SubsRPCPayload, resp *string) error {
	*resp = "ok"
	return ss.subsService.CreateFreeTrial(payload.UserId, payload.Kind)
}

func (ss *SubsServer) GetAccessTime(userId int, resp *int) error {
	subscription, err := ss.subsService.GetSubscription(userId)
	if err != nil {
		return err
	}
	*resp = subscription.AccessTime
	return nil
}
