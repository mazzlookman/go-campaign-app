// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependency_injection

import (
	"github.com/gin-gonic/gin"
	"go-campaign-app/app"
	"go-campaign-app/controller"
	"go-campaign-app/repository"
	"go-campaign-app/service"
)

// Injectors from injector.go:

func InitializedServer() *gin.Engine {
	campaignRepository := repository.NewCampaignRepository()
	db := app.DBConnection()
	campaignService := service.NewCampaignService(campaignRepository, db)
	campaignController := controller.NewCampaignController(campaignService)
	engine := app.NewRouter(campaignController)
	return engine
}

func InitializedServerTest() *gin.Engine {
	campaignRepository := repository.NewCampaignRepository()
	db := app.DBConnectionTest()
	campaignService := service.NewCampaignService(campaignRepository, db)
	campaignController := controller.NewCampaignController(campaignService)
	engine := app.NewRouter(campaignController)
	return engine
}