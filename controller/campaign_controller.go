package controller

import "github.com/gin-gonic/gin"

type CampaignController interface {
	RegisterUser(c *gin.Context)
}
