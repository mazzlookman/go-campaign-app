package repository

import (
	"go-campaign-app/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	//User Repository
	Save(db *gorm.DB, user domain.User) (domain.User, error)
	FindByEmail(db *gorm.DB, email string) (domain.User, error)
	FindById(db *gorm.DB, id int) (domain.User, error)
	UploadAvatar(db *gorm.DB, user domain.User) (domain.User, error)
}
