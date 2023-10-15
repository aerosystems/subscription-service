package RPCServer

type SubsRPCPayload struct {
	UserId uint
	Kind   string
}

func (ss *SubsServer) CreateFreeTrial(payload SubsRPCPayload, resp *string) error {
	*resp = "ok"
	return ss.subsService.CreateFreeTrial(payload.UserId, payload.Kind)
}

func (ss *SubsServer) GetAccessTime(userId uint, resp *uint) error {
	subscription, err := ss.subsService.GetSubscription(userId)
	if err != nil {
		return err
	}
	*resp = subscription.AccessTime
	return nil
}
