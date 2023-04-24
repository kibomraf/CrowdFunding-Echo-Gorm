package transaction

import "gorm.io/gorm"

type Repository interface {
	GetTransactionByCampaignId(campaignID int) ([]Transactions, error)
	GetByUserId(userID int) ([]Transactions, error)
	Save(t Transactions) (Transactions, error)
}
type repository struct {
	db *gorm.DB
}

func TransactionRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetTransactionByCampaignId(campaignID int) ([]Transactions, error) {
	var transaction []Transactions
	err := r.db.Where("campaign_id = ?", campaignID).Preload("User").Order("id desc").Find(&transaction).Error
	if err != nil {
		return []Transactions{}, err
	}
	return transaction, nil
}
func (r *repository) GetByUserId(userID int) ([]Transactions, error) {
	var transaction []Transactions
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = true").Where("user_id = ?", userID).Find(&transaction).Error
	if err != nil {
		return []Transactions{}, err
	}
	return transaction, nil
}
func (r *repository) Save(transaction Transactions) (Transactions, error) {
	err := r.db.Create(&transaction).Preload("Campaign").Error
	if err != nil {
		return Transactions{}, err
	}
	return transaction, nil
}
