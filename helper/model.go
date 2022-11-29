package helper

import (
	"go-campaign-app/model/domain"
	"go-campaign-app/model/web"
)

func UserFiltered(user *domain.User) web.UserFiltered {
	uf := web.UserFiltered{}
	uf.Id = user.Id
	uf.Name = user.Name
	uf.Occupation = user.Occupation
	uf.Email = user.Email
	uf.AvatarFileName = user.AvatarFileName.String
	uf.Role = user.Role

	return uf
}

func CampaignResponseFormatter(campaign *domain.Campaign) web.CampaignResponse {
	camp := web.CampaignResponse{}

	camp.Id = campaign.Id
	camp.UserId = campaign.UserId
	camp.Name = campaign.Name
	camp.ShortDescription = campaign.Summary
	if len(campaign.CampaignImages) > 0 {
		camp.ImageUrl = campaign.CampaignImages[0].FileName
	}
	camp.GoalAmount = campaign.GoalAmount
	camp.CurrentAmount = campaign.CurrentAmount

	return camp
}

func CampaignsResponseFormatter(campaigns []domain.Campaign) []web.CampaignResponse {
	camps := []web.CampaignResponse{}
	for _, c := range campaigns {
		campaignResponse := CampaignResponseFormatter(&c)
		camps = append(camps, campaignResponse)
	}
	return camps
}
