package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-campaign-app/helper"
	"go-campaign-app/model/web"
	"go-campaign-app/service"
	"net/http"
)

type CampaignControllerImpl struct {
	service.CampaignService
}

func (contr *CampaignControllerImpl) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		helper.ErrorUploadAvatar(err, c, http.StatusUnprocessableEntity)
		return
	}
	idUser := 1
	dst := fmt.Sprintf("images/%d-%s", idUser, file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		helper.ErrorUploadAvatar(err, c, http.StatusBadRequest)
		return
	}

	_, err = contr.CampaignService.UploadAvatar(c, dst, idUser)
	if err != nil {
		helper.ErrorCampaignService(err, c)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.WriteToResponseBody(200, "success", "Avatar successfully uploaded", data)
	c.JSON(200, &response)
}

func (contr *CampaignControllerImpl) CheckEmailAvailable(c *gin.Context) {
	email := web.CheckEmailAvailable{}
	err := c.ShouldBindJSON(&email)
	if err != nil {
		validationInput := helper.ErrorValidationInput(err, c)
		response := helper.WriteToResponseBody(http.StatusUnprocessableEntity, "error input", "Email checking is failed", validationInput)
		c.JSON(http.StatusUnprocessableEntity, &response)
		return
	}

	emailAvailable, err := contr.CampaignService.CheckEmailAvailable(c, email)
	if err != nil {
		helper.ErrorCampaignService(err, c)
	}

	message := "Email is available"
	if emailAvailable == false {
		message = "Email has been registered"
	}

	data := gin.H{
		"is_available": emailAvailable,
	}

	resp := helper.WriteToResponseBody(
		200,
		"success",
		message,
		data,
	)

	c.JSON(200, &resp)
}

func (contr *CampaignControllerImpl) RegisterUser(c *gin.Context) {
	user := web.RegisterUser{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		validationInput := helper.ErrorValidationInput(err, c)
		response := helper.WriteToResponseBody(http.StatusUnprocessableEntity, "error input", "Register account failed", validationInput)
		c.JSON(http.StatusUnprocessableEntity, &response)
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
		validationInput := helper.ErrorValidationInput(err, c)
		response := helper.WriteToResponseBody(http.StatusUnprocessableEntity, "error input", "Login is failed", validationInput)
		c.JSON(http.StatusUnprocessableEntity, &response)
		return
	}

	user, err := contr.CampaignService.LoginUser(c, login)
	if err != nil {
		helper.ErrorCampaignService(err, c)
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

func NewCampaignController(campaignService service.CampaignService) CampaignController {
	return &CampaignControllerImpl{CampaignService: campaignService}
}
