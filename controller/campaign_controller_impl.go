package controller

import (
	"github.com/gin-gonic/gin"
	"go-campaign-app/formatter"
	"go-campaign-app/helper"
	"go-campaign-app/model/web"
	"go-campaign-app/service"
	"net/http"
	"strconv"
)

type CampaignControllerImpl struct {
	service.CampaignService
}

func (contr *CampaignControllerImpl) UpdateCampaign(ctx *gin.Context) {
	inputId := web.FindCampaignById{}
	err := ctx.ShouldBindUri(&inputId)
	if err != nil {
		apiResponse := formatter.WriteToResponseBody(
			http.StatusUnprocessableEntity,
			"error",
			"couldn't mapping uri input to struct",
			err.Error(),
		)
		ctx.JSON(http.StatusUnprocessableEntity, &apiResponse)
	}

	camp := web.CreateCampaignInput{}
	err2 := ctx.ShouldBindJSON(&camp)
	if err2 != nil {
		apiResponse := formatter.WriteToResponseBody(
			http.StatusUnprocessableEntity,
			"error",
			"couldn't mapping json input to struct",
			err.Error(),
		)
		ctx.JSON(http.StatusUnprocessableEntity, &apiResponse)
		return
	}

	id := ctx.MustGet("currentUser").(int)
	camp.UserId = id

	updateCampaign, err3 := contr.CampaignService.UpdateCampaign(inputId, camp)
	if err3 != nil {
		apiResponse := formatter.WriteToResponseBody(
			http.StatusUnprocessableEntity,
			"error",
			"Failed to update campaign",
			err3.Error(),
		)
		ctx.JSON(http.StatusUnprocessableEntity, &apiResponse)
		return
	}

	apiResponse := formatter.WriteToResponseBody(
		200,
		"success",
		"Campaign successfully updated",
		formatter.CampaignResponseFormatter(&updateCampaign),
	)

	ctx.JSON(200, &apiResponse)
}

func (contr *CampaignControllerImpl) CreateCampaign(ctx *gin.Context) {
	input := web.CreateCampaignInput{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		apiResponse := formatter.WriteToResponseBody(
			http.StatusUnprocessableEntity,
			"error",
			"couldn't mapping json input to struct",
			err.Error(),
		)
		ctx.JSON(http.StatusUnprocessableEntity, &apiResponse)
	}

	userId := ctx.MustGet("currentUser").(int)
	input.UserId = userId

	createCampaign, err := contr.CampaignService.CreateCampaign(input)
	if err != nil {
		helper.CampaignServiceError(err)
		return
	}
	apiResponse := formatter.WriteToResponseBody(
		200,
		"success",
		"campaign has been created",
		formatter.CampaignResponseFormatter(&createCampaign),
	)
	ctx.JSON(200, &apiResponse)
}

func (contr *CampaignControllerImpl) FindCampaignById(ctx *gin.Context) {
	input := web.FindCampaignById{}
	err := ctx.ShouldBindUri(&input)
	if err != nil {
		apiResponse := formatter.WriteToResponseBody(
			http.StatusUnprocessableEntity,
			"error",
			"could't mapping uri input to struct",
			nil,
		)
		ctx.JSON(http.StatusUnprocessableEntity, &apiResponse)
		return
	}

	campaign, err := contr.CampaignService.FindCampaignById(input)
	if err != nil {
		helper.ErrorService(err, ctx)
		return
	}
	apiResponse := formatter.WriteToResponseBody(
		200,
		"success",
		"campaign detail",
		formatter.CampaignDetailFormatter(&campaign),
	)
	ctx.JSON(200, &apiResponse)
}

func (contr *CampaignControllerImpl) FindCampaigns(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	campaigns, err := contr.CampaignService.FindCampaigns(userId)
	if err != nil {
		helper.ErrorService(err, ctx)
		return
	}
	apiResponse := formatter.WriteToResponseBody(
		200,
		"success",
		"List of campaigns",
		formatter.CampaignsResponseFormatter(campaigns),
	)
	ctx.JSON(200, &apiResponse)
}

func NewCampaignControllerImpl(campaignService service.CampaignService) *CampaignControllerImpl {
	return &CampaignControllerImpl{CampaignService: campaignService}
}
