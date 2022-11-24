package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	CheckEmailAvailable(c *gin.Context)
	UploadAvatar(c *gin.Context)
}
