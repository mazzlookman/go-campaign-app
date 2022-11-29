package controller

import (
	"github.com/gin-gonic/gin"
	"go-campaign-app/helper"
	"go-campaign-app/service"
	"strconv"
)

type CampaignControllerImpl struct {
	service.CampaignService
}

func (contr *CampaignControllerImpl) FindCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := contr.CampaignService.FindCampaigns(userId)
	if err != nil {
		helper.ErrorService(err, c)
		return
	}
	apiResponse := helper.WriteToResponseBody(
		200,
		"success",
		"List of campaigns",
		helper.CampaignsResponseFormatter(campaigns),
	)
	c.JSON(200, &apiResponse)
}

func NewCampaignControllerImpl(campaignService service.CampaignService) *CampaignControllerImpl {
	return &CampaignControllerImpl{CampaignService: campaignService}
}
