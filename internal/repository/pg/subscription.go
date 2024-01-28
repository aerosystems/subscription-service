package pg

import (
	"errors"
	"github.com/aerosystems/subs-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) Create(subscription *models.Subscription) error {
	result := r.db.Create(&subscription)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *SubscriptionRepo) GetByUserUuid(userUuid uuid.UUID) (*models.Subscription, error) {
	var subscription models.Subscription
	result := r.db.Where("user_uuid = ?", userUuid.String()).First(&subscription)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &subscription, nil
}

func (r *SubscriptionRepo) Update(subscription *models.Subscription) error {
	result := r.db.Save(&subscription)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *SubscriptionRepo) Delete(subscription *models.Subscription) error {
	result := r.db.Delete(&subscription)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
