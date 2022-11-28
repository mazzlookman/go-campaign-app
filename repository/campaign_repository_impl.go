package repository

import (
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
}

func (repo *CampaignRepositoryImpl) FindAll(db *gorm.DB) ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := db.Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		helper.CampaignRepositoryError(err)
		return campaigns, err
	}
	return campaigns, nil
}

func (repo *CampaignRepositoryImpl) FindByUserId(db *gorm.DB, userId int) ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := db.Where("user_id=?", userId).
		Preload("CampaignImages", "campaign_images.is_primary=1").
		Find(&campaigns).
		Error

	if err != nil {
		helper.CampaignRepositoryError(err)
		return campaigns, err
	}

	return campaigns, nil
}

func NewCampaignRepository() CampaignRepository {
	return &CampaignRepositoryImpl{}
}
