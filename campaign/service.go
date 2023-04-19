package campaign

import (
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaign(userId int) ([]Campaigns, error)
	GetDetailsCampaign(input InputGetDetailCampaign) (Campaigns, error)
	CreateCampaign(input CreateCampaignInput) (Campaigns, error)
}
type campaignService struct {
	r Repository
}

func Campaignservices(r Repository) *campaignService {
	return &campaignService{r}
}
func (s *campaignService) GetCampaign(userId int) ([]Campaigns, error) {
	if userId != 0 {
		campaign, err := s.r.FindUserById(userId)
		if err != nil {
			return campaign, err
		}
		return campaign, nil
	}
	campaign, err := s.r.FindAll()
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s *campaignService) GetDetailsCampaign(input InputGetDetailCampaign) (Campaigns, error) {
	campaigns, err := s.r.FindById(input.Id)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil

}
func (s *campaignService) CreateCampaign(input CreateCampaignInput) (Campaigns, error) {
	campaign := Campaigns{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       uint(input.GoalAmount),
		Perks:            input.Perks,
		User_id:          input.User.Id,
		Created_at:       time.Now(),
		Updated_at:       time.Now(),
	}
	slugCandidate := fmt.Sprintf("%v %v", input.Name, input.User.Id)
	campaign.Slug = slug.Make(slugCandidate)
	//pembuatan slug
	newCampaign, err := s.r.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
