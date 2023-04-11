package campaign

type Service interface {
	GetCampaign(userId int) ([]Campaigns, error)
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
