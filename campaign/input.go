package campaign

import "crowdfunding/users"

type InputGetDetailCampaign struct {
	Id int `json:"id"`
}
type CreateCampaignInput struct {
	Name             string      `json:"name" validate:"required"`
	ShortDescription string      `json:"short_description" validate:"required"`
	Description      string      `json:"description" validate:"required"`
	GoalAmount       int         `json:"goal_amount" validate:"required"`
	Perks            string      `json:"perks" validate:"required"`
	User             users.Users `json:"user" validate:"required"`
}
