package repository

import (
	"context"
	"database/sql"
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
)

type UserRepositoryImpl struct {
}

func (repo *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
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

func (repo *UserRepositoryImpl) UploadAvatar(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	sql := "update users set avatar = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, user.AvatarFileName.String, user.Id)
	helper.PanicIfError(err)

	return user, nil
}

func (repo *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	sql := "insert into users (name,occupation,email,password_hash,avatar,role) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, user.Name, user.Occupation, user.Email, user.PasswordHash, user.AvatarFileName, user.Role)
	if err != nil {
		return user, err
	}
	id, _ := result.LastInsertId()
	user.Id = int(id)

	return user, nil
}

func (repo *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
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

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
