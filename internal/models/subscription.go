package models

type Subscription struct {
	Id         int64            `json:"id"`
	Kind       KindSubscription `json:"kind"`
	UserId     int64            `json:"userId"`
	AccessTime int64            `json:"accessTime"`
	CreatedAt  int64            `json:"createdAt"`
	UpdatedAt  int64            `json:"updatedAt"`
}

type KindSubscription string

const (
	Startup  KindSubscription = "startup"
	Business KindSubscription = "business"
)

type SubscriptionRepository interface {
	Create(subscription *Subscription) error
	GetByUserId(userId int64) ([]*Subscription, error)
	Update(subscription *Subscription) error
	Delete(id int64) error
}
