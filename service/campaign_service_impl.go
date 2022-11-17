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

type CampaignServiceImpl struct {
	repository.CampaignRepository
	*sql.DB
}

func (service *CampaignServiceImpl) CheckEmailAvailable(ctx context.Context, available web.CheckEmailAvailable) (bool, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := available.Email

	user, err := service.CampaignRepository.FindByEmail(ctx, tx, email)
	helper.PanicIfError(err)

	if user.Id == 0 {
		return true, nil
	}
	return false, nil
}

func (service *CampaignServiceImpl) RegisterUser(ctx context.Context, user web.RegisterUser) (web.UserFiltered, error) {
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

	save, err := service.CampaignRepository.Save(ctx, tx, userRepo)
	if err != nil {
		return helper.UserFiltered(&save), err
	}

	return helper.UserFiltered(&save), nil
}

func (service *CampaignServiceImpl) LoginUser(ctx context.Context, user web.LoginUser) (web.UserFiltered, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := user.Email
	pass := user.Password

	findByEmail, err := service.CampaignRepository.FindByEmail(ctx, tx, email)
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

func NewCampaignService(campaignRepository repository.CampaignRepository, DB *sql.DB) CampaignService {
	return &CampaignServiceImpl{CampaignRepository: campaignRepository, DB: DB}
}
