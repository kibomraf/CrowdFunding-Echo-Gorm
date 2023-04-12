package campaign

type FormatCampaign struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short description"`
	GoalAmount       int    `json:"goal amount"`
	CurrentAmount    int    `json:"current amount"`
	Slug             string `json:"slug"`
	ImageUrl         string `json:"image url"`
}

func FormatterCampaign(campaign Campaigns) FormatCampaign {
	formatCampaign := FormatCampaign{
		Id:               campaign.Id,
		UserId:           campaign.User_id,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       int(campaign.GoalAmount),
		CurrentAmount:    int(campaign.CurrentAmount),
		Slug:             campaign.Slug,
		ImageUrl:         "",
	}
	if len(campaign.CampaignImages) > 0 {
		formatCampaign.ImageUrl = campaign.CampaignImages[0].FileName
	}
	return formatCampaign
}
func FormatterCampaigns(campaign []Campaigns) []FormatCampaign {
	formatCampaigns := []FormatCampaign{}
	for _, campaign := range campaign {
		formatCampaigns = append(formatCampaigns, FormatterCampaign(campaign))
	}
	return formatCampaigns
}
