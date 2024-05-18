package pg

import (
	"context"
	"errors"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

type SubscriptionPg struct {
	Id         int       `gorm:"primaryKey;autoIncrement"`
	UserUuid   uuid.UUID `gorm:"unique"`
	Kind       string    `gorm:"<-"`
	Duration   string    `gorm:"<-"`
	AccessTime time.Time `gorm:"<-"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (r *SubscriptionPg) ToModel() *models.Subscription {
	return &models.Subscription{
		UserUuid:   r.UserUuid,
		Type:       models.NewSubscriptionType(r.Kind),
		Duration:   models.NewSubscriptionDuration(r.Duration),
		AccessTime: r.AccessTime,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

func ModelToPg(subscription *models.Subscription) *SubscriptionPg {
	return &SubscriptionPg{
		UserUuid:   subscription.UserUuid,
		Kind:       subscription.Type.String(),
		Duration:   subscription.Duration.String(),
		AccessTime: subscription.AccessTime,
		CreatedAt:  subscription.CreatedAt,
		UpdatedAt:  subscription.UpdatedAt,
	}
}

func (r *SubscriptionRepo) Create(ctx context.Context, subscription *models.Subscription) error {
	result := r.db.Create(ModelToPg(subscription))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *SubscriptionRepo) GetByUserUuid(ctx context.Context, userUuid uuid.UUID) (*models.Subscription, error) {
	var subscription SubscriptionPg
	result := r.db.Where("user_uuid = ?", userUuid.String()).First(&subscription)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return subscription.ToModel(), nil
}

func (r *SubscriptionRepo) Update(ctx context.Context, subscription *models.Subscription) error {
	result := r.db.Save(ModelToPg(subscription))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *SubscriptionRepo) Delete(ctx context.Context, subscription *models.Subscription) error {
	result := r.db.Delete(ModelToPg(subscription))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
