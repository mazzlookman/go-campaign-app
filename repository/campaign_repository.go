package repository

import (
	"go-campaign-app/model/domain"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll(db *gorm.DB) ([]domain.Campaign, error)
	FindByUserId(db *gorm.DB, userId int) ([]domain.Campaign, error)
}
