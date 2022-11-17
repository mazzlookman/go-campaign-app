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

func ErrorValidationInput(err error, c *gin.Context) {
	var errors []string

	for _, fieldError := range err.(validator.ValidationErrors) {
		errors = append(errors, fieldError.Error())
	}

	errorsMap := gin.H{"errors": errors}
	response := WriteToResponseBody(http.StatusUnprocessableEntity, "error input", "Register account failed", errorsMap)
	c.JSON(http.StatusUnprocessableEntity, &response)
}

func ErrorCampaignService(err error, c *gin.Context) {
	response := WriteToResponseBody(http.StatusInternalServerError, "INTERNAL SERVER ERROR", "Register account failed", err.Error())
	c.JSON(http.StatusInternalServerError, &response)
}
