package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-campaign-app/app"
	"go-campaign-app/helper"
	"go-campaign-app/model/web"
	"go-campaign-app/repository"
	"go-campaign-app/service"
	"testing"
)

func TestCampaignRepositoryFindById(t *testing.T) {
	d := app.DBConnection()
	r := repository.NewCampaignRepository()
	campaigns, err := r.FindByUserId(d, 1)
	helper.PanicIfError(err)
	fmt.Println(len(campaigns))
	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
		if len(campaign.CampaignImages) > 0 {
			fmt.Println(campaign.CampaignImages[0].FileName)
		}
		fmt.Println("==================")
	}
}

func TestCampaignServiceFind(t *testing.T) {
	s := service.NewCampaignService(repository.NewCampaignRepository(), app.DBConnection())
	campaigns, err := s.FindCampaigns(0)
	helper.PanicIfError(err)
	assert.Equal(t, 3, len(campaigns))

	user1, _ := s.FindCampaigns(1)
	assert.Equal(t, 2, len(user1))

	user2, _ := s.FindCampaigns(2)
	assert.Equal(t, 1, len(user2))
}

func TestCreateCampaignService(t *testing.T) {
	serv := service.NewUserService(repository.NewUserRepository(), app.DBConnection())
	user, _ := serv.FindById(3)

	campService := service.NewCampaignService(repository.NewCampaignRepository(), app.DBConnection())
	input := web.CreateCampaignInput{
		Name:             "Penggalangan Dana Startup",
		ShortDescription: "pds",
		Description:      "pds lagi",
		GoalAmount:       150000000,
		Perks:            "perks 1, perks 2, perks 3",
		UserId:           user.Id,
	}
	createCampaign, err := campService.CreateCampaign(input)
	helper.CampaignServiceError(err)
	fmt.Println(createCampaign)
}
