//go:build wireinject
// +build wireinject

package dependency_injection

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-campaign-app/app"
	"go-campaign-app/controller"
	"go-campaign-app/repository"
	"go-campaign-app/service"
)

//
//func InitializedCampaignServiceTest() *service.CampaignServiceImpl {
//	wire.Build(
//		app.DBConnectionTest,
//		repository.NewCampaignRepository,
//		service.NewCampaignService,
//	)
//	return nil
//}
//
////func InitializedCampaignControllerTest() *controller.CampaignControllerImpl {
////	wire.Build(
////		app.DBConnectionTest,
////		repository.NewCampaignRepository,
////		service.NewCampaignService,
////		controller.NewCampaignController,
////	)
////	return nil
////}

func InitializedServer() *gin.Engine {
	wire.Build(
		app.DBConnection,
		repository.NewCampaignRepository,
		service.NewCampaignService,
		controller.NewCampaignController,
		app.NewRouter,
	)
	return nil
}

func InitializedServerTest() *gin.Engine {
	wire.Build(
		app.DBConnectionTest,
		repository.NewCampaignRepository,
		service.NewCampaignService,
		controller.NewCampaignController,
		app.NewRouter,
	)
	return nil
}
