package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"crowdfunding/helper"
	"crowdfunding/transaction"
	"crowdfunding/users"
)

type transactionHandler struct {
	service transaction.Service
}

func TransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}
func (h *transactionHandler) GetTransactionByCampaignId(c echo.Context) error {
	var parameter transaction.GetTransactionByCampaignIdInput
	id, err := strconv.Atoi(c.Param("id"))
	parameter.Id = id
	if err != nil {
		response := helper.APIResponse("error to get data transactions by campaign id", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	validation := validator.New()
	err = validation.Struct(parameter)
	if err != nil {
		response := helper.APIResponse("error to get parameter", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	transactions, err := h.service.GetTransactionByCampaignId(parameter)
	if err != nil {
		response := helper.APIResponse("error to get data transactions by campaign id", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	formatter := transaction.GetTransactionByCampaignIdFormatters(transactions)
	response := helper.APIResponse("success to get campaign's transactions", http.StatusOK, "error", formatter)
	return c.JSON(echo.ErrBadRequest.Code, response)
}
func (h *transactionHandler) GetTransactionOfUser(c echo.Context) error {
	currentUser := c.Get("user").(users.Users)
	var input transaction.GetTransactionByUserIdInput
	input.Id = currentUser.Id
	err := c.Bind(input)
	if err != nil {
		response := helper.APIResponse("error to get data user's transactions 1", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	validation := validator.New()
	err = validation.Struct(input)
	if err != nil {
		response := helper.APIResponse("error to get data user's transactions 2", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	transactions, err := h.service.GetTransactionByUserId(input)
	if err != nil {
		response := helper.APIResponse("error to get data user's transactions 3", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	formatter := transaction.UserTransactionsFormatter(transactions)
	response := helper.APIResponse("success to get user's transactions 4", http.StatusOK, "successfully", formatter)
	return c.JSON(echo.ErrBadRequest.Code, response)
}
func (h *transactionHandler) CreateTransaction(c echo.Context) error {
	var input transaction.CreateTransactionInput
	currentUser := c.Get("user").(users.Users)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("error to create transaction 1", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	input.CampaignId = id
	input.UserId = currentUser.Id
	err = c.Bind(&input)
	if err != nil {
		response := helper.APIResponse("error to get input parameter 2", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	validation := validator.New()
	err = validation.Struct(input)
	if err != nil {
		errorMsg := helper.FormatError(err)
		response := helper.APIResponse("unprocessable entity", echo.ErrUnprocessableEntity.Code, "unprocessable", errorMsg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	//call service to save
	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("error to create transaction 3", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	formatter := transaction.CreateTransactionFormater(newTransaction)
	response := helper.APIResponse("success to create transaction", http.StatusOK, "successfully", formatter)
	return c.JSON(echo.ErrBadRequest.Code, response)

}
