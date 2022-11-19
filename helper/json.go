package helper

import (
	"go-campaign-app/model/web"
)

func WriteToResponseBody(code int, st string, msg string, data interface{}) web.APIResponse {
	meta := web.Meta{
		Code:    code,
		Status:  st,
		Message: msg,
	}

	resp := web.APIResponse{
		Meta: meta,
		Data: data,
	}

	return resp
}

func UserResponseAPI(user web.UserFiltered, token string) web.RegisterOrLoginUserResponse {
	resp := web.RegisterOrLoginUserResponse{
		Id:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return resp
}
