package service

import (
	"context"
	"go-campaign-app/model/web"
)

type UserService interface {
	RegisterUser(ctx context.Context, user web.RegisterUser) (web.UserFiltered, error)
	LoginUser(ctx context.Context, user web.LoginUser) (web.UserFiltered, error)
	CheckEmailAvailable(ctx context.Context, available web.CheckEmailAvailable) (bool, error)
	UploadAvatar(ctx context.Context, fileName string, id int) (web.UserFiltered, error)
	FindById(ctx context.Context, id int) (web.UserFiltered, error)
}
