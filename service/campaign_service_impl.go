package service

import (
	"context"
	"database/sql"
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

func NewCampaignService(campaignRepository repository.CampaignRepository, DB *sql.DB) CampaignService {
	return &CampaignServiceImpl{CampaignRepository: campaignRepository, DB: DB}
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
