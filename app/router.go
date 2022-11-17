package app

import (
	"github.com/gin-gonic/gin"
	"go-campaign-app/controller"
)

func NewRouter(controller controller.CampaignController) *gin.Engine {
	router := gin.Default()
	group := router.Group("/api/v1")
	group.POST("/users", controller.RegisterUser)
	group.POST("/sessions", controller.LoginUser)

	return router
}
