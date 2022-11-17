package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorValidationInput(err error, c *gin.Context) *gin.H {
	var errors []string

	for _, fieldError := range err.(validator.ValidationErrors) {
		errors = append(errors, fieldError.Error())
	}

	errorsMap := gin.H{"errors": errors}

	return &errorsMap
}

func ErrorCampaignService(err error, c *gin.Context) {
	response := WriteToResponseBody(http.StatusInternalServerError, "INTERNAL SERVER ERROR", "Ups sorry", err.Error())
	c.JSON(http.StatusInternalServerError, &response)
}
