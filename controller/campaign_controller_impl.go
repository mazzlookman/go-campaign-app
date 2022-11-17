package controller

import (
	"github.com/gin-gonic/gin"
	"go-campaign-app/helper"
	"go-campaign-app/model/web"
	"go-campaign-app/service"
	"net/http"
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

func (contr *CampaignControllerImpl) LoginUser(c *gin.Context) {
	login := web.LoginUser{}

	err := c.ShouldBindJSON(&login)
	if err != nil {
		helper.ErrorValidationInput(err, c)
		return
	}

	user, err := contr.CampaignService.LoginUser(c, login)
	if err != nil {
		response := helper.WriteToResponseBody(http.StatusBadRequest, "BAD REQUEST", "Login is failed", err.Error())
		c.JSON(http.StatusInternalServerError, &response)
		return
	}

	userResponse := helper.UserResponseAPI(user, "tokentokentoken")
	response := helper.WriteToResponseBody(
		200,
		"success",
		"Login successfully",
		userResponse)

	c.JSON(200, response)
}
