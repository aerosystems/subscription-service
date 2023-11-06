package repository

import (
	"github.com/aerosystems/subs-service/internal/models"
	"gorm.io/gorm"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) SubscriptionRepo {
	return SubscriptionRepo{db: db}
}

func (r SubscriptionRepo) Create(subscription *models.Subscription) error {
	return r.db.Create(subscription).Error
}

func (r SubscriptionRepo) GetByUserId(userId int) (*models.Subscription, error) {
	var subscriptions *models.Subscription
	err := r.db.Where("user_id = ?", userId).Find(&subscriptions).Error
	return subscriptions, err
}

func (r SubscriptionRepo) Update(subscription *models.Subscription) error {
	return r.db.Save(subscription).Error
}

func (r SubscriptionRepo) Delete(id int) error {
	return r.db.Delete(&models.Subscription{}, id).Error
}
