package transaction

import (
	"time"

	"crowdfunding/campaign"
	"crowdfunding/users"
)

type Transactions struct {
	ID              int
	CampaignId      int
	UserId          int
	Amount          uint
	Status          string
	CodeTransaction string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Campaign        campaign.Campaigns
	User            users.Users
}
