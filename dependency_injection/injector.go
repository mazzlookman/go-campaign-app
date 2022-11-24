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
		repository.NewUserRepository,
		service.NewUserService,
		middleware.NewJWTAuthImpl,
		controller.NewUserController,
		app.NewRouter,
	)
	return nil
}

func InitializedServerTest() *gin.Engine {
	wire.Build(
		app.DBConnectionTest,
		repository.NewUserRepository,
		service.NewUserService,
		middleware.NewJWTAuthImpl,
		controller.NewUserController,
		app.NewRouter,
	)
	return nil
}

func InitializedJwtAuthMiddleware() gin.HandlerFunc {
	wire.Build(
		app.DBConnectionTest,
		repository.NewUserRepository,
		service.NewUserService,
		middleware.NewJWTAuthImpl,
		middleware.NewJWTAuthMiddleware,
	)
	return nil
}
