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
