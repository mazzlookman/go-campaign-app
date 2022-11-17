package repository

import (
	"context"
	"database/sql"
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
)

type CampaignRepositoryImpl struct {
}

func NewCampaignRepository() CampaignRepository {
	return &CampaignRepositoryImpl{}
}

func (repo *CampaignRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	sql := "insert into users (name,occupation,email,password_hash,avatar,role) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, user.Name, user.Occupation, user.Email, user.PasswordHash, user.AvatarFileName, user.Role)
	if err != nil {
		return user, err
	}
	id, _ := result.LastInsertId()
	user.Id = int(id)

	return user, nil
}

func (repo *CampaignRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	sql := "select * from users where email = ?"
	rows, err := tx.QueryContext(ctx, sql, email)
	if err != nil {
		helper.PanicIfError(err)
	}
	user := domain.User{}
	if rows.Next() {
		rows.Scan(
			&user.Id, &user.Name, &user.Occupation, &user.Email, &user.PasswordHash, &user.AvatarFileName,
			&user.Role, &user.CreatedAt, &user.UpdatedAt,
		)
	}
	return user, nil
}

