package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-campaign-app/app"
	"go-campaign-app/controller"
	"go-campaign-app/helper"
	"go-campaign-app/middleware"
	"go-campaign-app/model/domain"
	"go-campaign-app/model/web"
	"go-campaign-app/repository"
	"go-campaign-app/service"
	"io"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
)

var db = app.DBConnectionTest()
var ctx = context.Background()
var repo = repository.NewUserRepository()
var serv = service.NewUserService(repo, db)
var contr = controller.NewUserController(serv, middleware.NewJWTAuthImpl())

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

func TestCampaignRepositoryFindByEmail(t *testing.T) {
	tx, _ := db.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := repo.FindByEmail(ctx, tx, "uup@test.com")
	helper.PanicIfError(err)

	fmt.Println(user)
	assert.Equal(t, "", user.Occupation)
}

func TestCampaignServiceLoginUser(t *testing.T) {
	tx, _ := db.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := serv.LoginUser(ctx, web.LoginUser{
		Email:    "ucup@test.com",
		Password: "password",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, 5, user.Id)
}

func TestServiceCheckEmailAvailable(t *testing.T) {
	emailAvailable, err := serv.CheckEmailAvailable(ctx, web.CheckEmailAvailable{Email: "up@test.com"})
	helper.PanicIfError(err)

	assert.Equal(t, false, emailAvailable)
}

func TestNameServiceUpdateAvatar(t *testing.T) {
	fileName := "image/avatar6.jpg"
	user, err := serv.UploadAvatar(ctx, fileName, 1)
	helper.PanicIfError(err)

	fmt.Println(user)
	assert.Equal(t, 1, user.Id)
	assert.Equal(t, fileName, user.AvatarFileName)
}

func TestGenerateToken(t *testing.T) {
	var m = map[string]string{
		"email":    "aqib@test.com",
		"password": "password",
	}
	marshal, _ := json.Marshal(&m)

	requ := httptest.NewRequest("POST", "http://localhost:2802/api/v1/sessions", strings.NewReader(string(marshal)))
	writer := httptest.NewRecorder()

	router := app.NewRouter(contr)
	router.ServeHTTP(writer, requ)

	response := writer.Result()
	bytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

func TestValidateToken(t *testing.T) {
	jwtAuth := middleware.NewJWTAuthImpl()
	token, err := jwtAuth.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.PnCL06cxJiB1R3cg17EAJAcXDnFwPU2QSp3lubIyQ_o")
	helper.PanicIfError(err)

	if token.Valid {
		log.Println("token is validated")
	} else {
		log.Println("token not valid " + err.Error())
	}
}
