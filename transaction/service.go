package transaction

import "time"

type Service interface {
	GetTransactionByCampaignId(campaign GetTransactionByCampaignIdInput) ([]Transactions, error)
	GetTransactionByUserId(user GetTransactionByUserIdInput) ([]Transactions, error)
	CreateTransaction(input CreateTransactionInput) (Transactions, error)
}
type service struct {
	repo Repository
}

func TransactionService(repo Repository) *service {
	return &service{repo}
}
func (s *service) GetTransactionByCampaignId(campaign GetTransactionByCampaignIdInput) ([]Transactions, error) {
	transaction, err := s.repo.GetTransactionByCampaignId(campaign.Id)
	if err != nil {
		return []Transactions{}, err
	}
	return transaction, nil
}
func (s *service) GetTransactionByUserId(user GetTransactionByUserIdInput) ([]Transactions, error) {
	transactions, err := s.repo.GetByUserId(user.Id)
	if err != nil {
		return []Transactions{}, err
	}
	return transactions, nil
}
func (s *service) CreateTransaction(input CreateTransactionInput) (Transactions, error) {
	transaction := Transactions{
		CampaignId:      input.CampaignId,
		UserId:          input.UserId,
		Amount:          input.Amount,
		Status:          "pending",
		CodeTransaction: "KIBO-",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	newTransaction, err := s.repo.Save(transaction)
	if err != nil {
		return Transactions{}, err
	}
	return newTransaction, nil
}
