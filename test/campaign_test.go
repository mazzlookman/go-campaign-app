package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-campaign-app/app"
	"go-campaign-app/helper"
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
	campaigns, err := s.Find(0)
	helper.PanicIfError(err)
	assert.Equal(t, 3, len(campaigns))

	user1, _ := s.Find(1)
	assert.Equal(t, 2, len(user1))

	user2, _ := s.Find(2)
	assert.Equal(t, 1, len(user2))
}
