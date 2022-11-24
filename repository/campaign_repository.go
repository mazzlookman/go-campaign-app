package repository

import (
	"context"
	"database/sql"
	"go-campaign-app/model/domain"
)

type CampaignRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Campaign, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) (domain.Campaign, error)
}
