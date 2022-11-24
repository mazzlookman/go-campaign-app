package test

import (
	"context"
	"database/sql"
	"fmt"
	"go-campaign-app/app"
	"go-campaign-app/helper"
	"go-campaign-app/repository"
	"testing"
)

func dbContext() (*sql.DB, context.Context) {
	d := app.DBConnection()
	c := context.Background()
	return d, c
}

func TestFindAllCampaign(t *testing.T) {
	d, c := dbContext()
	tx, _ := d.Begin()
	r := repository.NewCampaignRepository()
	campaigns, err := r.FindAll(c, tx)
	helper.PanicIfError(err)
	fmt.Println(campaigns)
}
