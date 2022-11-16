package main

import (
	"github.com/gin-gonic/gin"
	"go-campaign-app/test"
)

func main() {
	router := gin.Default()
	router.GET("/handler", test.HandlerTest)
	router.Run()
}
