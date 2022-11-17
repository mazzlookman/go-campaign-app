package test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-campaign-app/app"
	"go-campaign-app/controller"
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
	"go-campaign-app/model/web"
	"go-campaign-app/repository"
	"go-campaign-app/service"
	"testing"
)

var db = app.DBConnectionTest()
var ctx = context.Background()
var repo = repository.NewCampaignRepository()
var serv = service.NewCampaignService(repo, db)
var contr = controller.NewCampaignController(serv)

//var addr = "http://localhost:8080/api/v1"

func TestDBConnection(t *testing.T) {
	tx, err := db.Begin()
	helper.PanicIfError(err)

	result, err := tx.Query("select * from users")
	helper.PanicIfError(err)
	defer result.Close()

	var users []domain.User
	for result.Next() {
		var user domain.User
		result.Scan(
			&user.Id,
			&user.Name,
			&user.Occupation,
			&user.Email,
			&user.PasswordHash,
			&user.AvatarFileName,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		users = append(users, user)
	}

	for _, user := range users {
		fmt.Println(user)
	}

	tx.Commit()
}

func TestCampaignRepositorySave(t *testing.T) {
	tx, _ := db.Begin()
	defer helper.CommitOrRollback(tx)

	var user domain.User
	user.Name = "Otong"

	save, err := repo.Save(ctx, tx, user)
	helper.PanicIfError(err)

	assert.Equal(t, user.Name, save.Name)
}

func TestCampaignServiceRegisterUser(t *testing.T) {
	user := web.RegisterUser{
		Name:       "Ucup",
		Occupation: "UI/UX",
		Email:      "ucup@test.com",
		Password:   "password",
	}

	registerUser, err := serv.RegisterUser(ctx, user)
	helper.PanicIfError(err)

	assert.Equal(t, user.Email, registerUser.Email)
}

//func TestRegisterUserController(t *testing.T) {
//	payload := strings.NewReader(`{"name":"Teguh","occupation": "Data Analyst","email":"teguh@test.com","password":"password"}`)
//	request := httptest.NewRequest("POST", addr+"/users", payload)
//
//	writer := httptest.NewRecorder()
//
//	http.Handler.ServeHTTP(contr.RegisterUser)
//}
