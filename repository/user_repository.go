package repository

import (
	"context"
	"database/sql"
	"go-campaign-app/model/domain"
)

type UserRepository interface {
	//User Repository
	Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	UploadAvatar(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)

	//Campaign Repository

}
