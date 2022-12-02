package web

type CampaignResponse struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type CampaignDetailFormatter struct {
	Id               int                             `json:"id"`
	Name             string                          `json:"name"`
	ShortDescription string                          `json:"short_description"`
	Description      string                          `json:"description"`
	ImageUrl         string                          `json:"image_url"`
	GoalAmount       int                             `json:"goal_amount"`
	CurrentAmount    int                             `json:"current_amount"`
	UserId           int                             `json:"user_id"`
	Slug             string                          `json:"slug"`
	Perks            []string                        `json:"perks"`
	User             CampaignDetailUserFormatter     `json:"user"`
	Images           []CampaignDetailImagesFormatter `json:"images"`
}

type CampaignDetailUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignDetailImagesFormatter struct {
	FileName  string `json:"file_name"`
	IsPrimary bool   `json:"is_primary"`
}
