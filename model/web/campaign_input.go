package web

type FindCampaignById struct {
	Id int `uri:"campaignId" binding:"required"`
}
