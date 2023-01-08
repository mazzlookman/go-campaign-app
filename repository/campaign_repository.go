package repository

import (
	"go-campaign-app/model/domain"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll(db *gorm.DB) ([]domain.Campaign, error)
	FindByUserId(db *gorm.DB, userId int) ([]domain.Campaign, error)
	FindById(db *gorm.DB, id int) (domain.Campaign, error)
	Save(db *gorm.DB, campaign domain.Campaign) (domain.Campaign, error)
	Update(db *gorm.DB, campaign domain.Campaign) (domain.Campaign, error)
}
