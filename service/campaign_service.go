package service

import (
	"context"
	"go-campaign-app/model/web"
)

type CampaignService interface {
	RegisterUser(ctx context.Context, user web.RegisterUser) (web.UserFiltered, error)
}