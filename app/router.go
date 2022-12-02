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
	router.Static("/images", "./images")

	//User endpoint
	api := router.Group("/api/v1")
	api.POST("/users", contrUser.RegisterUser)
	api.POST("/sessions", contrUser.LoginUser)
	api.POST("/email-checker", contrUser.CheckEmailAvailable)
	api.POST("/avatars", jwtAuthMiddleware, contrUser.UploadAvatar)

	//Campaign endpoint
	api.GET("/campaigns", contrCampaign.FindCampaigns)
	api.GET("/campaigns/:campaignId", contrCampaign.FindCampaignById)

	return router
}
