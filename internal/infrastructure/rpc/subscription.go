package RPCServer

import (
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
	"time"
)

type SubsRPCPayload struct {
	UserUuid   uuid.UUID
	Kind       models.KindSubscription
	AccessTime time.Time
}

func (s Server) CreateFreeTrial(payload SubsRPCPayload, resp *string) error {
	*resp = "ok"
	return s.subscriptionUsecase.CreateFreeTrial(payload.UserUuid, payload.Kind)
}

func (s Server) GetSubscription(userUuid uuid.UUID, resp *SubsRPCPayload) error {
	sub, err := s.subscriptionUsecase.GetSubscription(userUuid)
	if err != nil {
		return err
	}
	resp.Kind = sub.Kind
	resp.AccessTime = sub.AccessTime
	return nil
}

func (s Server) DeleteSubscription(userUuid uuid.UUID, resp *string) error {
	*resp = "ok"
	return s.subscriptionUsecase.DeleteSubscription(userUuid)
}
