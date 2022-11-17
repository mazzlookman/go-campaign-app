package controller

import (
	"github.com/gin-gonic/gin"
	"go-campaign-app/helper"
	"go-campaign-app/model/web"
	"go-campaign-app/service"
)

type CampaignControllerImpl struct {
	service.CampaignService
}

func NewCampaignController(campaignService service.CampaignService) CampaignController {
	return &CampaignControllerImpl{CampaignService: campaignService}
}

func (contr *CampaignControllerImpl) RegisterUser(c *gin.Context) {
	user := web.RegisterUser{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		helper.ErrorValidationInput(err, c)
		return
	}

	registerUser, err := contr.CampaignService.RegisterUser(c, user)
	if err != nil {
		helper.ErrorCampaignService(err, c)
		return
	}

	userResponse := helper.UserResponseAPI(registerUser, "tokentokentokentoken")

	apiResponse := helper.WriteToResponseBody(
		200,
		"success",
		"Account has been registered",
		userResponse)

	c.JSON(200, &apiResponse)
}
