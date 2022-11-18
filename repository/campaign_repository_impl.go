package repository

import (
	"context"
	"database/sql"
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
)

type CampaignRepositoryImpl struct {
}

func (repo *CampaignRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	sql := "select * from users where id = ?"
	rows, err := tx.QueryContext(ctx, sql, id)
	defer rows.Close()
	helper.PanicIfError(err)

	user := domain.User{}
	if rows.Next() {
		rows.Scan(
			&user.Id, &user.Name, &user.Occupation, &user.Email, &user.PasswordHash, &user.AvatarFileName,
			&user.Role, &user.CreatedAt, &user.UpdatedAt,
		)
	}

	return user, nil
}

func (repo *CampaignRepositoryImpl) UpdateAvatar(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	sql := "update users set avatar = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, user.AvatarFileName, user.Id)
	helper.PanicIfError(err)

	return user, nil
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
	defer rows.Close()
	helper.PanicIfError(err)

	user := domain.User{}
	if rows.Next() {
		rows.Scan(
			&user.Id, &user.Name, &user.Occupation, &user.Email, &user.PasswordHash, &user.AvatarFileName,
			&user.Role, &user.CreatedAt, &user.UpdatedAt,
		)
	}

	return user, nil
}

func NewCampaignRepository() CampaignRepository {
	return &CampaignRepositoryImpl{}
}
