package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaigns, error)
	FindUserById(userId int) ([]Campaigns, error)
}
type repository struct {
	db *gorm.DB
}

func CampaignRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindAll() ([]Campaigns, error) {
	var campaigns []Campaigns
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
func (r *repository) FindUserById(userId int) ([]Campaigns, error) {
	var campaigns []Campaigns
	err := r.db.Where("id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
