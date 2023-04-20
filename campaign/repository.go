package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaigns, error)
	FindUserById(userId int) ([]Campaigns, error)
	FindById(userId int) (Campaigns, error)
	Save(camp Campaigns) (Campaigns, error)
	Update(camp Campaigns) (Campaigns, error)
	SaveImages(image CampaignImages) (CampaignImages, error)
	MarkAllImagesAsNonPrimary(campID int) (bool, error)
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
func (r *repository) FindById(userId int) (Campaigns, error) {
	var campaign Campaigns
	err := r.db.Preload("CampaignImages").Preload("User").Where("id = ?", userId).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
func (r *repository) Save(camp Campaigns) (Campaigns, error) {
	err := r.db.Create(&camp).Error
	if err != nil {
		return camp, err
	}
	return camp, nil
}
func (r *repository) Update(camp Campaigns) (Campaigns, error) {
	err := r.db.Save(&camp).Error
	if err != nil {
		return camp, err
	}
	return camp, nil
}
func (r *repository) SaveImages(image CampaignImages) (CampaignImages, error) {
	err := r.db.Create(&image).Error
	if err != nil {
		return image, err
	}
	return image, nil
}
func (r *repository) MarkAllImagesAsNonPrimary(campID int) (bool, error) {
	err := r.db.Model(&CampaignImages{}).Where("campaign_id = ?", campID).Update("is_primary", false).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}
