package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"crowdfunding/campaign"
	"crowdfunding/helper"
)

type campHandler struct {
	sevice campaign.Service
}

func NewCampHandler(service campaign.Service) *campHandler {
	return &campHandler{service}
}
func (h *campHandler) GetCampaignc(c echo.Context) error {
	userId, _ := strconv.Atoi(c.QueryParam("user_id"))
	camp, err := h.sevice.GetCampaign(userId)
	if err != nil {
		response := helper.APIResponse("error", echo.ErrInternalServerError.Code, "error", nil)
		return c.JSON(echo.ErrInternalServerError.Code, response)
	}
	formatter := campaign.FormatterCampaigns(camp)
	response := helper.APIResponse("list of campaigns", http.StatusOK, "success", formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *campHandler) GetDetailsCampaign(c echo.Context) error {
	var input campaign.InputGetDetailCampaign
	input.Id, _ = strconv.Atoi(c.Param("id"))
	campaigns, err := h.sevice.GetDetailsCampaign(input)
	if err != nil {
		response := helper.APIResponse("error", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	formatter := campaign.GetDetailsCampaignFormatter(campaigns)
	response := helper.APIResponse("details of campaign", http.StatusOK, "successfully", formatter)
	return c.JSON(http.StatusOK, response)
}

//tangkap parameter user ke input struct
//panggil service parameternya input
//panggil repository untuk simpan data campain baru
