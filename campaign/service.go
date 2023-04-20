package campaign

import (
	"errors"
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaign(userId int) ([]Campaigns, error)
	GetDetailsCampaign(input InputGetDetailCampaign) (Campaigns, error)
	CreateCampaign(input CreateCampaignInput) (Campaigns, error)
	UpdateCampaign(campId InputGetDetailCampaign, update CreateCampaignInput) (Campaigns, error)
	SaveImagesCampaign(input InputSaveImages, fileLocation string) (CampaignImages, error)
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
	newCampaign, err := s.r.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
func (s *campaignService) UpdateCampaign(campId InputGetDetailCampaign, update CreateCampaignInput) (Campaigns, error) {
	campaign, err := s.r.FindById(campId.Id)
	if err != nil {
		return campaign, err
	}
	if update.User.Id != campaign.User_id {
		msg := errors.New("you are not owner the campaign, do you want to iDOR ?")
		return campaign, msg
	}
	campaign.Name = update.Name
	campaign.ShortDescription = update.ShortDescription
	campaign.Description = update.Description
	campaign.Perks = update.Perks
	campaign.GoalAmount = uint(update.GoalAmount)
	updateCampaign, err := s.r.Update(campaign)
	if err != nil {
		return campaign, err
	}
	return updateCampaign, nil
}
func (s *campaignService) SaveImagesCampaign(input InputSaveImages, fileLocation string) (CampaignImages, error) {
	campaign, err := s.r.FindById(input.CampaignID)
	if err != nil {
		return CampaignImages{}, err
	}
	if campaign.User_id != input.User.Id {
		return CampaignImages{}, errors.New("you're not the owner")
	}
	var isPrimary bool
	if input.IsPrimary == "true" {
		isPrimary = true
		_, err := s.r.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImages{}, err
		}
	}
	campaignImage := CampaignImages{
		CampaignId: input.CampaignID,
		FileName:   fileLocation,
		IsPrimary:  isPrimary,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	newImages, err := s.r.SaveImages(campaignImage)
	if err != nil {
		return CampaignImages{}, nil
	}
	return newImages, nil
}
