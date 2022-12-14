package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-campaign-app/formatter"
	"net/http"
)

func UserRepositoryError(err error) {
	if err != nil {
		panic(errors.New("Ups. User Repository Error " + err.Error()))
	}
}

func CampaignRepositoryError(err error) {
	if err != nil {
		panic(errors.New("Ups. Campaign Repository Error " + err.Error()))
	}
}

func UserServiceError(err error) {
	if err != nil {
		panic(errors.New("Ups. User Service Error " + err.Error()))
	}
}

func CampaignServiceError(err error) {
	if err != nil {
		panic(errors.New("Ups. Campaign Service Error " + err.Error()))
	}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorValidationInput(err error) *gin.H {
	var errors []string

	for _, fieldError := range err.(validator.ValidationErrors) {
		errors = append(errors, fieldError.Error())
	}

	errorsMap := gin.H{"errors": errors}

	return &errorsMap
}

func ErrorService(err error, c *gin.Context) {
	response := formatter.WriteToResponseBody(http.StatusInternalServerError, "INTERNAL SERVER ERROR", "Ups sorry", err.Error())
	c.JSON(http.StatusInternalServerError, &response)
}

func ErrorUploadAvatar(err error, c *gin.Context, code int) {
	data := gin.H{"is_uploaded": false}
	response := formatter.WriteToResponseBody(code, "error", "Upload avatar is failed", data)
	c.JSON(code, &response)
}

func ErrorUnauthorized(c *gin.Context, code int) {
	response := formatter.WriteToResponseBody(code, "error", "UNAUTHORIZED", nil)
	c.AbortWithStatusJSON(code, &response)
}
