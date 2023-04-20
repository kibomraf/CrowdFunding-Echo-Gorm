package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"crowdfunding/campaign"
	"crowdfunding/helper"
	"crowdfunding/users"
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

func (h *campHandler) CreateCampaign(c echo.Context) error {
	var input campaign.CreateCampaignInput
	err := c.Bind(&input)
	if err != nil {
		response := helper.APIResponse("failed to create campaign", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMsg := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("unprocessable entity", echo.ErrUnprocessableEntity.Code, "error", errorMsg)
		return c.JSON(echo.ErrUnprocessableEntity.Code, response)
	}
	currentUser := c.Get("user").(users.Users)
	input.User = currentUser
	newCampaign, err := h.sevice.CreateCampaign(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMsg := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("failed to create campaign", echo.ErrBadRequest.Code, "error", errorMsg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	formatter := campaign.FormatterCampaign(newCampaign)
	response := helper.APIResponse("success to create campaign", http.StatusOK, "success", formatter)
	return c.JSON(http.StatusOK, response)
}
func (h *campHandler) UpdateCampaign(c echo.Context) error {
	var inputID campaign.InputGetDetailCampaign
	var err error
	inputID.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		errors := helper.FormatError(err)
		errorMsg := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("failed to create campaign", echo.ErrBadRequest.Code, "error", errorMsg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	var camp campaign.CreateCampaignInput
	err = c.Bind(&camp)
	if err != nil {
		errors := helper.FormatError(err)
		errorMsg := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("failed to create campaign", echo.ErrBadRequest.Code, "error", errorMsg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	validate := validator.New()
	err = validate.Struct(camp)
	if err != nil {
		errors := helper.FormatError(err)
		errorMsg := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("unprocessable entity", echo.ErrUnprocessableEntity.Code, "error", errorMsg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	currentUser := c.Get("user").(users.Users)
	camp.User = currentUser
	updatedCampaign, err := h.sevice.UpdateCampaign(inputID, camp)
	if err != nil {
		errors := helper.FormatError(err)
		errorMsg := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("failed to create campaign", echo.ErrBadRequest.Code, "error", errorMsg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	formatter := campaign.FormatterCampaign(updatedCampaign)
	response := helper.APIResponse("success to updated campaign", http.StatusOK, "successfully", formatter)
	return c.JSON(http.StatusOK, response)

}

// update campaign
// user masukan input
// handler
// mapping dari input ke input struc
// input user dan juga input dari uri
// service (find campaignbyid, tankap parameter)
// repository update campaign
