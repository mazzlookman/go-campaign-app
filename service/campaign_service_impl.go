package service

import (
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
	"go-campaign-app/repository"
	"gorm.io/gorm"
)

type CampaignServiceImpl struct {
	repository.CampaignRepository
	*gorm.DB
}

func (c *CampaignServiceImpl) FindCampaigns(userId int) ([]domain.Campaign, error) {
	if userId != 0 {
		campaigns, err := c.CampaignRepository.FindByUserId(c.DB, userId)
		helper.CampaignServiceError(err)

		return campaigns, nil
	}

	campaigns, err := c.CampaignRepository.FindAll(c.DB)
	helper.CampaignServiceError(err)

	return campaigns, nil
}

func NewCampaignService(campaignRepository repository.CampaignRepository, DB *gorm.DB) CampaignService {
	return &CampaignServiceImpl{CampaignRepository: campaignRepository, DB: DB}
}
