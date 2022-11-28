package test

import (
	"fmt"
	"go-campaign-app/app"
	"go-campaign-app/helper"
	"go-campaign-app/repository"
	"testing"
)

func TestFindAllCampaign(t *testing.T) {
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
