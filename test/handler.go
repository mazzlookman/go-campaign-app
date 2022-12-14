package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-campaign-app/app"
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
)

func HandlerTest(c *gin.Context) {
	db := app.DBConnectionTest()
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

	c.JSON(200, &users)
}
