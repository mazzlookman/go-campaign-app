package formatter

import (
	"go-campaign-app/model/domain"
	"go-campaign-app/model/web"
)

func UserFiltered(user *domain.User) web.UserFiltered {
	uf := web.UserFiltered{}
	uf.Id = user.Id
	uf.Name = user.Name
	uf.Occupation = user.Occupation
	uf.Email = user.Email
	uf.AvatarFileName = user.AvatarFileName
	uf.Role = user.Role

	return uf
}

func UserRegisterOrLogin(user web.UserFiltered, token string) web.RegisterOrLoginUserResponse {
	resp := web.RegisterOrLoginUserResponse{
		Id:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return resp
}
