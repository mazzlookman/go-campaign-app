package repository

import (
	"context"
	"database/sql"
	"go-campaign-app/model/domain"
)

type CampaignRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
}
