package app

import (
	"github.com/gin-gonic/gin"
	"go-campaign-app/controller"
	"go-campaign-app/middleware"
	"go-campaign-app/repository"
	"go-campaign-app/service"
)

func NewRouter() *gin.Engine {
	db := DBConnectionTest()
	repo := repository.NewCampaignRepository()
	serv := service.NewCampaignService(repo, db)
	jwtAuth := middleware.NewJWTAuthImpl()
	contr := controller.NewCampaignController(serv, jwtAuth)
	jwtAuthMiddleware := middleware.NewJWTAuthMiddleware(jwtAuth, serv)

	router := gin.Default()
	group := router.Group("/api/v1")
	group.POST("/users", contr.RegisterUser)
	group.POST("/sessions", contr.LoginUser)
	group.POST("/email-checker", contr.CheckEmailAvailable)
	group.POST("/avatars", jwtAuthMiddleware, contr.UploadAvatar)

	return router
}
