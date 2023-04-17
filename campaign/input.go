package campaign

import "crowdfunding/users"

type InputGetDetailCampaign struct {
	Id int `json:"id"`
}
type CreateCampaignInput struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	GoalAmount       int    `json:"goal_amount"`
	Perks            string `json:"perks"`
	User             users.Users
}
