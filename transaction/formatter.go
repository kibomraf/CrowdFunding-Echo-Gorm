package transaction

import "time"

type GetTransactionCampaignFormat struct {
	Id        int    `json:"id"`
	UserName  string `json:"user_name"`
	Status    string
	Amount    uint      `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
type TransactionUserFormat struct {
	Id        int       `json:"id"`
	Campaign  Campaign  `json:"campaign"`
	Amount    uint      `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
type Campaign struct {
	CampaignName string `json:"campaign_name"`
	Image        string `json:"image"`
}
type CreateTransactionFormat struct {
	CampaignId int       `json:"campaign_id"`
	Amount     uint      `json:"amount"`
	Status     string    `json:"status"`
	CreateAt   time.Time `json:"created_at"`
}

func GetTransactionByCampaignIdFormatter(t Transactions) GetTransactionCampaignFormat {
	format := GetTransactionCampaignFormat{
		Id:        t.ID,
		UserName:  t.User.Name,
		Amount:    t.Amount,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
	}
	return format
}
func GetTransactionByCampaignIdFormatters(t []Transactions) []GetTransactionCampaignFormat {
	formatters := []GetTransactionCampaignFormat{}
	if len(t) == 0 {
		return []GetTransactionCampaignFormat{}
	}
	for _, transactions := range t {
		format := GetTransactionByCampaignIdFormatter(transactions)
		if format.Status == "paid" {
			formatters = append(formatters, format)
		}
	}
	return formatters
}
func UserTransactionFormatter(t Transactions) TransactionUserFormat {
	format := TransactionUserFormat{
		Id: t.ID,
		Campaign: Campaign{
			CampaignName: t.Campaign.Name,
			Image:        t.Campaign.CampaignImages[0].FileName,
		},
		Amount:    t.Amount,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
	}
	return format
}
func UserTransactionsFormatter(t []Transactions) []TransactionUserFormat {
	if len(t) == 0 {
		return []TransactionUserFormat{}
	}
	formatter := []TransactionUserFormat{}
	for _, transaction := range t {
		format := UserTransactionFormatter(transaction)
		formatter = append(formatter, format)
	}
	return formatter
}
func CreateTransactionFormater(input Transactions) CreateTransactionFormat {
	formatter := CreateTransactionFormat{
		CampaignId: input.CampaignId,
		Amount:     input.Amount,
		Status:     input.Status,
		CreateAt:   input.CreatedAt,
	}
	return formatter
}
