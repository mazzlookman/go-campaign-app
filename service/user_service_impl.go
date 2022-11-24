package service

import (
	"context"
	"database/sql"
	"errors"
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
	"go-campaign-app/model/web"
	"go-campaign-app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository.UserRepository
	*sql.DB
}

func (service *UserServiceImpl) FindById(ctx context.Context, id int) (web.UserFiltered, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.UserFiltered(&user), nil
}

func (service *UserServiceImpl) UploadAvatar(ctx context.Context, fileName string, id int) (web.UserFiltered, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	user.AvatarFileName.String = fileName

	avatar, err := service.UserRepository.UploadAvatar(ctx, tx, user)
	helper.PanicIfError(err)

	return helper.UserFiltered(&avatar), nil
}

func (service *UserServiceImpl) CheckEmailAvailable(ctx context.Context, available web.CheckEmailAvailable) (bool, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := available.Email

	user, err := service.UserRepository.FindByEmail(ctx, tx, email)
	helper.PanicIfError(err)

	if user.Id == 0 {
		return true, nil
	}
	return false, nil
}

func (service *UserServiceImpl) RegisterUser(ctx context.Context, user web.RegisterUser) (web.UserFiltered, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRepo := domain.User{}
	userRepo.Name = user.Name
	userRepo.Occupation = user.Occupation
	userRepo.Email = user.Email
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	helper.PanicIfError(err)
	userRepo.PasswordHash = string(password)
	userRepo.Role = "user"

	save, err := service.UserRepository.Save(ctx, tx, userRepo)
	if err != nil {
		return helper.UserFiltered(&save), err
	}

	return helper.UserFiltered(&save), nil
}

func (service *UserServiceImpl) LoginUser(ctx context.Context, user web.LoginUser) (web.UserFiltered, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := user.Email
	pass := user.Password

	findByEmail, err := service.UserRepository.FindByEmail(ctx, tx, email)
	helper.PanicIfError(err)

	if findByEmail.Id == 0 {
		return web.UserFiltered{}, errors.New("Account not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.PasswordHash), []byte(pass))
	if err != nil {
		return web.UserFiltered{}, errors.New("Your password is incorrect")
	}

	return helper.UserFiltered(&findByEmail), nil
}

func NewUserService(campaignRepository repository.UserRepository, DB *sql.DB) UserService {
	return &UserServiceImpl{UserRepository: campaignRepository, DB: DB}
}
