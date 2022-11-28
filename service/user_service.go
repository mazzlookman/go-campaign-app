package service

import (
	"go-campaign-app/model/web"
)

type UserService interface {
	RegisterUser(user web.RegisterUser) (web.UserFiltered, error)
	LoginUser(user web.LoginUser) (web.UserFiltered, error)
	CheckEmailAvailable(available web.CheckEmailAvailable) (bool, error)
	UploadAvatar(fileName string, id int) (web.UserFiltered, error)
	FindById(id int) (web.UserFiltered, error)
}
