package formatter

import (
	"go-campaign-app/model/domain"
	"go-campaign-app/model/web"
	"strings"
)

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
	camp.Slug = campaign.Slug

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

func CampaignDetailFormatter(campaign *domain.Campaign) web.CampaignDetailFormatter {
	camp := web.CampaignDetailFormatter{}
	camp.Id = campaign.Id
	camp.Name = campaign.Name
	camp.ShortDescription = campaign.Summary
	camp.Description = campaign.Description

	//image_url
	if len(campaign.CampaignImages) > 0 {
		for _, image := range campaign.CampaignImages {
			if image.IsPrimary == 1 {
				camp.ImageUrl = image.FileName
			}
		}
	}

	camp.GoalAmount = campaign.GoalAmount
	camp.CurrentAmount = campaign.CurrentAmount
	camp.UserId = campaign.UserId
	camp.Slug = campaign.Slug

	//Perks
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	camp.Perks = perks

	//User
	user := campaign.User
	userDetail := web.CampaignDetailUserFormatter{
		Name:     user.Name,
		ImageUrl: user.AvatarFileName,
	}
	camp.User = userDetail

	//CampaignImages
	campaignImages := []web.CampaignDetailImagesFormatter{}
	images := campaign.CampaignImages
	for _, image := range images {
		campImage := web.CampaignDetailImagesFormatter{}
		campImage.FileName = image.FileName
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
			campImage.IsPrimary = isPrimary
		}
		campaignImages = append(campaignImages, campImage)
	}
	camp.Images = campaignImages

	return camp
}
