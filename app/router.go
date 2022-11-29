package app

import (
	"github.com/gin-gonic/gin"
	"go-campaign-app/controller"
	"go-campaign-app/middleware"
	"go-campaign-app/repository"
	"go-campaign-app/service"
)

func NewRouter() *gin.Engine {
	db := DBConnection()

	repoUser := repository.NewUserRepository()
	repoCampaign := repository.NewCampaignRepository()

	servUser := service.NewUserService(repoUser, db)
	servCampaign := service.NewCampaignService(repoCampaign, db)

	jwtAuth := middleware.NewJWTAuthImpl()
	jwtAuthMiddleware := middleware.NewJWTAuthMiddleware(jwtAuth, servUser)

	contrUser := controller.NewUserController(servUser, jwtAuth)
	contrCampaign := controller.NewCampaignControllerImpl(servCampaign)

	router := gin.Default()
	//User endpoint
	group := router.Group("/api/v1")
	group.POST("/users", contrUser.RegisterUser)
	group.POST("/sessions", contrUser.LoginUser)
	group.POST("/email-checker", contrUser.CheckEmailAvailable)
	group.POST("/avatars", jwtAuthMiddleware, contrUser.UploadAvatar)

	//Campaign endpoint
	group.GET("/campaigns", contrCampaign.FindCampaigns)

	return router
}
