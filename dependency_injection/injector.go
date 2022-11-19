//go:build wireinject
// +build wireinject

package dependency_injection

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-campaign-app/app"
	"go-campaign-app/controller"
	"go-campaign-app/middleware"
	"go-campaign-app/repository"
	"go-campaign-app/service"
)

func InitializedServer() *gin.Engine {
	wire.Build(
		app.DBConnection,
		repository.NewCampaignRepository,
		service.NewCampaignService,
		middleware.NewJWTAuthImpl,
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
		middleware.NewJWTAuthImpl,
		controller.NewCampaignController,
		app.NewRouter,
	)
	return nil
}

func InitializedJwtAuthMiddleware() gin.HandlerFunc {
	wire.Build(
		app.DBConnectionTest,
		repository.NewCampaignRepository,
		service.NewCampaignService,
		middleware.NewJWTAuthImpl,
		middleware.NewJWTAuthMiddleware,
	)
	return nil
}
