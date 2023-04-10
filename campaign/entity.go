package campaign

import "time"

type Campaigns struct {
	Id               int
	User_id          int
	Name             string
	Backer_account   int
	ShortDescription string
	GoalAmount       uint
	CurrentAmount    uint
	Perks            string
	Slug             string
	Created_at       time.Time
	Updated_at       time.Time
	CampaignImages   []CampaignImages `gorm:"foreignKey:campaign_id"`
}
type CampaignImages struct {
	Id         int
	CampaignId int
	FileName   string
	IsPrimary  bool
	Created_at time.Time
	Updated_at time.Time
}
