package repository

import (
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
}

func (repo *CampaignRepositoryImpl) Save(db *gorm.DB, campaign domain.Campaign) (domain.Campaign, error) {
	err := db.Create(&campaign).Error
	helper.CampaignRepositoryError(err)

	return campaign, nil
}

func (repo *CampaignRepositoryImpl) FindById(db *gorm.DB, id int) (domain.Campaign, error) {
	camp := domain.Campaign{}
	err := db.Preload("User").
		Preload("CampaignImages").
		Where("id=?", id).
		Find(&camp).
		Error
	helper.CampaignRepositoryError(err)

	return camp, nil
}

func (repo *CampaignRepositoryImpl) FindAll(db *gorm.DB) ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := db.Preload("CampaignImages", "campaign_images.is_primary=1").
		Find(&campaigns).
		Error
	helper.CampaignRepositoryError(err)

	return campaigns, nil
}

func (repo *CampaignRepositoryImpl) FindByUserId(db *gorm.DB, userId int) ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := db.Preload("CampaignImages", "campaign_images.is_primary=1").
		Where("user_id=?", userId).
		Find(&campaigns).
		Error

	helper.CampaignRepositoryError(err)

	return campaigns, nil
}

func NewCampaignRepository() CampaignRepository {
	return &CampaignRepositoryImpl{}
}
