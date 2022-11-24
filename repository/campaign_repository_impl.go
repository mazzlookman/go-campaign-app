package repository

import (
	"context"
	"database/sql"
	"go-campaign-app/helper"
	"go-campaign-app/model/domain"
)

type CampaignRepositoryImpl struct {
}

func (repo *CampaignRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Campaign, error) {
	sql := "select * from campaigns"
	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		helper.CampaignRepositoryError(err)
	}
	defer rows.Close()

	var campaigns []domain.Campaign
	for rows.Next() {
		campaign := domain.Campaign{}
		rows.Scan(&campaign.Id, &campaign.UserId, &campaign.Name, &campaign.Summary, &campaign.Description,
			&campaign.Perks, &campaign.BackerCount, &campaign.GoalAmount, &campaign.CurrentAmount,
			&campaign.Slug, &campaign.CreatedAt, &campaign.UpdatedAt)

		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}

func (repo *CampaignRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) (domain.Campaign, error) {
	sql := "select * from campaigns where id = ?"
	rows, err := tx.QueryContext(ctx, sql, userId)
	if err != nil {
		helper.CampaignRepositoryError(err)
	}
	defer rows.Close()

	var campaign domain.Campaign
	if rows.Next() {
		rows.Scan(&campaign.Id, &campaign.UserId, &campaign.Name, &campaign.Summary, &campaign.Description,
			&campaign.Perks, &campaign.BackerCount, &campaign.GoalAmount, &campaign.CurrentAmount,
			&campaign.Slug, &campaign.CreatedAt, &campaign.UpdatedAt)
	}

	return campaign, nil
}

func NewCampaignRepository() CampaignRepository {
	return &CampaignRepositoryImpl{}
}
