package repository

import (
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func (repo *UserRepositoryImpl) FindById(db *gorm.DB, id int) (domain.User, error) {
	user := domain.User{}
	err := db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		helper.UserRepositoryError(err)
		return user, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) UploadAvatar(db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.Save(&user).Where("id", user.Id).Error
	if err != nil {
		helper.UserRepositoryError(err)
		return user, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) Save(db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		helper.UserRepositoryError(err)
		return user, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) FindByEmail(db *gorm.DB, email string) (domain.User, error) {
	user := domain.User{}
	err := db.Where("email=?", email).Find(&user).Error
	if err != nil {
		helper.UserRepositoryError(err)
		return user, err
	}

	return user, nil
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
