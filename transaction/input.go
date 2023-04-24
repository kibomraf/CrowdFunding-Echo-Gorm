package transaction

type GetTransactionByCampaignIdInput struct {
	Id int `url:"id" binding:"required"`
}
type GetTransactionByUserIdInput struct {
	Id int
}
type CreateTransactionInput struct {
	CampaignId int `url:"id" binding:"required"`
	UserId     int
	Amount     uint `json:"amount" validate:"required"`
}
