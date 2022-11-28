package service

import (
	"errors"
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
	"go-campaign-app/model/web"
	"go-campaign-app/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	repository.UserRepository
	*gorm.DB
}

func (service *UserServiceImpl) FindById(id int) (web.UserFiltered, error) {
	user, err := service.UserRepository.FindById(service.DB, id)
	helper.UserServiceError(err)

	return helper.UserFiltered(&user), nil
}

func (service *UserServiceImpl) UploadAvatar(fileName string, id int) (web.UserFiltered, error) {
	user, err := service.UserRepository.FindById(service.DB, id)
	helper.UserServiceError(err)

	user.AvatarFileName.String = fileName

	avatar, err := service.UserRepository.UploadAvatar(service.DB, user)
	helper.UserServiceError(err)

	return helper.UserFiltered(&avatar), nil
}

func (service *UserServiceImpl) CheckEmailAvailable(available web.CheckEmailAvailable) (bool, error) {
	email := available.Email

	user, err := service.UserRepository.FindByEmail(service.DB, email)
	helper.UserServiceError(err)

	if user.Id == 0 {
		return true, nil
	}
	return false, nil
}

func (service *UserServiceImpl) RegisterUser(user web.RegisterUser) (web.UserFiltered, error) {
	userRepo := domain.User{}
	userRepo.Name = user.Name
	userRepo.Occupation = user.Occupation
	userRepo.Email = user.Email
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	helper.PanicIfError(err)
	userRepo.PasswordHash = string(password)
	userRepo.Role = "user"
	save, err := service.UserRepository.Save(service.DB, userRepo)
	helper.UserServiceError(err)

	return helper.UserFiltered(&save), nil
}

func (service *UserServiceImpl) LoginUser(user web.LoginUser) (web.UserFiltered, error) {
	email := user.Email
	pass := user.Password

	findByEmail, err := service.UserRepository.FindByEmail(service.DB, email)
	helper.UserServiceError(err)

	if findByEmail.Id == 0 {
		return web.UserFiltered{}, errors.New("Account not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.PasswordHash), []byte(pass))
	if err != nil {
		return web.UserFiltered{}, errors.New("Your password is incorrect")
	}

	return helper.UserFiltered(&findByEmail), nil
}

func NewUserService(campaignRepository repository.UserRepository, DB *gorm.DB) UserService {
	return &UserServiceImpl{UserRepository: campaignRepository, DB: DB}
}
