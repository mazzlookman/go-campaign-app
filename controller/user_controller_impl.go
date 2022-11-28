package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-campaign-app/helper"
	"go-campaign-app/middleware"
	"go-campaign-app/model/web"
	"go-campaign-app/service"
	"net/http"
)

type UserControllerImpl struct {
	service.UserService
	middleware.JWTAuth
}

func (contr *UserControllerImpl) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		helper.ErrorUploadAvatar(err, c, http.StatusUnprocessableEntity)
		return
	}
	idUser := c.MustGet("currentUser").(int)

	dst := fmt.Sprintf("images/%d-%s", idUser, file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		helper.ErrorUploadAvatar(err, c, http.StatusBadRequest)
		return
	}

	_, err = contr.UserService.UploadAvatar(dst, idUser)
	if err != nil {
		helper.ErrorUserService(err, c)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.WriteToResponseBody(200, "success", "Avatar successfully uploaded", data)
	c.JSON(200, &response)
}

func (contr *UserControllerImpl) CheckEmailAvailable(c *gin.Context) {
	email := web.CheckEmailAvailable{}
	err := c.ShouldBindJSON(&email)
	if err != nil {
		validationInput := helper.ErrorValidationInput(err, c)
		response := helper.WriteToResponseBody(http.StatusUnprocessableEntity, "error input", "Email checking is failed", validationInput)
		c.JSON(http.StatusUnprocessableEntity, &response)
		return
	}

	emailAvailable, err := contr.UserService.CheckEmailAvailable(email)
	if err != nil {
		helper.ErrorUserService(err, c)
	}

	message := "Email is available"
	if emailAvailable == false {
		message = "Email has been registered"
	}

	data := gin.H{
		"is_available": emailAvailable,
	}

	resp := helper.WriteToResponseBody(
		200,
		"success",
		message,
		data,
	)

	c.JSON(200, &resp)
}

func (contr *UserControllerImpl) RegisterUser(c *gin.Context) {
	user := web.RegisterUser{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		validationInput := helper.ErrorValidationInput(err, c)
		response := helper.WriteToResponseBody(http.StatusUnprocessableEntity, "error input", "Register account failed", validationInput)
		c.JSON(http.StatusUnprocessableEntity, &response)
		return
	}

	registerUser, err := contr.UserService.RegisterUser(user)
	if err != nil {
		helper.ErrorUserService(err, c)
		return
	}

	token, err := contr.JWTAuth.GenerateToken(registerUser.Id)
	helper.PanicIfError(err)

	userResponse := helper.UserResponseAPI(registerUser, token)

	apiResponse := helper.WriteToResponseBody(
		200,
		"success",
		"Account has been registered",
		userResponse)

	c.JSON(200, &apiResponse)
}

func (contr *UserControllerImpl) LoginUser(c *gin.Context) {
	login := web.LoginUser{}

	err := c.ShouldBindJSON(&login)
	if err != nil {
		validationInput := helper.ErrorValidationInput(err, c)
		response := helper.WriteToResponseBody(http.StatusUnprocessableEntity, "error input", "Login is failed", validationInput)
		c.JSON(http.StatusUnprocessableEntity, &response)
		return
	}

	user, err := contr.UserService.LoginUser(login)
	if err != nil {
		helper.ErrorUserService(err, c)
		return
	}

	token, err := contr.JWTAuth.GenerateToken(user.Id)
	helper.PanicIfError(err)

	userResponse := helper.UserResponseAPI(user, token)
	response := helper.WriteToResponseBody(
		200,
		"success",
		"Login successfully",
		userResponse)

	c.JSON(200, response)
}

func NewUserController(userService service.UserService, auth middleware.JWTAuth) UserController {
	return &UserControllerImpl{
		UserService: userService,
		JWTAuth:     auth,
	}
}
