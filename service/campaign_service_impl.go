package service

import (
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
	"go-campaign-app/model/web"
	"go-campaign-app/repository"
	"gorm.io/gorm"
)

type CampaignServiceImpl struct {
	repository.CampaignRepository
	*gorm.DB
}

func (c *CampaignServiceImpl) FindCampaignById(id web.FindCampaignById) (domain.Campaign, error) {
	campaign, err := c.CampaignRepository.FindById(c.DB, id.Id)
	helper.CampaignServiceError(err)

	return campaign, nil
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
