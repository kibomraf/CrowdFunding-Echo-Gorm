package campaign

import "strings"

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

type CampaignDetailsFormatter struct {
	Id               int              `json:"id"`
	Name             string           `json:"name"`
	BackerAccount    int              `json:"backer_account"`
	ShortDescription string           `json:"short_description"`
	Description      string           `json:"description"`
	ImageUrl         string           `json:"image_url"`
	GoalAmount       int              `json:"goal_amount"`
	CurrentAmount    int              `json:"current_amount"`
	Slug             string           `json:"slug"`
	Perks            []string         `json:"perks"`
	User             UserCampaign     `json:"user"`
	Images           []ImagesCampaign `json:"images"`
}
type UserCampaign struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}
type ImagesCampaign struct {
	Name      string `json:"name"`
	IsPrimary bool   `json:"is_primary"`
}

func GetDetailsCampaignFormatter(camp Campaigns) CampaignDetailsFormatter {
	formatCampaign := CampaignDetailsFormatter{
		Id:               camp.Id,
		Name:             camp.Name,
		BackerAccount:    camp.Backer_account,
		ShortDescription: camp.ShortDescription,
		Description:      camp.Description,
		ImageUrl:         "",
		GoalAmount:       int(camp.GoalAmount),
		CurrentAmount:    int(camp.CurrentAmount),
		Slug:             camp.Slug,
		User: UserCampaign{
			Name:     camp.User.Name,
			ImageUrl: camp.User.Avatar_file_name,
		},
	}
	if len(camp.CampaignImages) > 0 {
		formatCampaign.ImageUrl = camp.CampaignImages[0].FileName
	}
	var perks []string
	for _, perk := range strings.Split(camp.Perks, ",") {
		perks = append(perks, perk)
	}
	formatCampaign.Perks = perks
	var images []ImagesCampaign
	for _, img := range camp.CampaignImages {
		images = append(images, ImagesCampaign{
			Name:      img.FileName,
			IsPrimary: img.IsPrimary,
		})
	}
	formatCampaign.Images = images
	return formatCampaign
}
