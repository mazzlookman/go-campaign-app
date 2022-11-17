package helper

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
