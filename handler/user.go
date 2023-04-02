package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"crowdfunding/helper"
	"crowdfunding/users"
)

type handler struct {
	service users.Service
}

func UserHandler(service users.Service) *handler {
	return &handler{service}
}
func (h *handler) RegisterUser(c echo.Context) error {
	var input users.InputRegister
	err := c.Bind(&input)
	if err != nil {
		response := helper.APIResponse("error request", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	// Validasi input register menggunakan package validator
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("unprocessable entity", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	user, err := h.service.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("error register user", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	// Format data user menggunakan formatter
	formatter := users.FormatterUsers(user, "token")

	// Set response API
	response := helper.APIResponse("success register user", http.StatusCreated, "success", formatter)
	return c.JSON(http.StatusCreated, response)
}

func (h *handler) LoginUser(c echo.Context) error {
	var login users.InputLogin
	err := c.Bind(&login)
	if err != nil {
		response := helper.APIResponse("error request", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	validate := validator.New()
	err = validate.Struct(login)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("unprocessable entity", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	loginUser, err := h.service.LoginUser(login)
	if err != nil {
		response := helper.APIResponse("not found", http.StatusNotFound, "email or password is invalid", nil)
		return c.JSON(http.StatusNotFound, response)
	}
	formater := users.FormatterUsers(loginUser, "token")
	response := helper.APIResponse("error request", http.StatusBadRequest, "error", formater)
	return c.JSON(http.StatusBadRequest, response)
}

func (h *handler) FetchUser(c echo.Context) error {
	currentuser := c.Get("currentUser").(users.Users)
	formatter := users.FormatterUsers(currentuser, "token")
	response := helper.APIResponse("success fetch user data", http.StatusOK, "success", formatter)
	return c.JSON(http.StatusOK, response)
}
