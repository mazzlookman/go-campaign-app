package service

import (
	"go-campaign-app/model/domain"
	"go-campaign-app/model/web"
)

type CampaignService interface {
	FindCampaigns(userId int) ([]domain.Campaign, error)
	FindCampaignById(id web.FindCampaignById) (domain.Campaign, error)
	CreateCampaign(input web.CreateCampaignInput) (domain.Campaign, error)
}
