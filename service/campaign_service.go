package service

import "go-campaign-app/model/domain"

type CampaignService interface {
	FindCampaigns(userId int) ([]domain.Campaign, error)
}
