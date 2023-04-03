package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
		response := helper.APIResponse("error request", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}

	// Validasi input register menggunakan package validator
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("unprocessable entity", echo.ErrUnprocessableEntity.Code, "error", errorMessage)
		return c.JSON(echo.ErrUnprocessableEntity.Code, response)
	}
	user, err := h.service.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("error register user", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
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
		response := helper.APIResponse("error request", echo.ErrBadRequest.Code, "error", nil)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	validate := validator.New()
	err = validate.Struct(login)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("unprocessable entity", echo.ErrUnprocessableEntity.Code, "error", errorMessage)
		return c.JSON(echo.ErrUnprocessableEntity.Code, response)
	}
	loginUser, err := h.service.LoginUser(login)
	if err != nil {
		response := helper.APIResponse("email or password is invalid", echo.ErrNotFound.Code, "not found", nil)
		return c.JSON(echo.ErrNotFound.Code, response)
	}
	formater := users.FormatterUsers(loginUser, "token")
	response := helper.APIResponse("login success", http.StatusOK, "succes", formater)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) FetchUser(c echo.Context) error {
	currentuser := c.Get("currentUser").(users.Users)
	formatter := users.FormatterUsers(currentuser, "token")
	response := helper.APIResponse("success fetch user data", http.StatusOK, "success", formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) CheckEmailAvailablity(c echo.Context) error {
	//input user terhadap email register
	//input map di mapping ke struct input
	//struct input di mapping ke service
	//service manggil repo mengencek terhadap email
	//repo - db
	var checkemail users.ChekEmail
	err := c.Bind(&checkemail)
	if err != nil {
		errors := helper.FormatError(err)
		msgError := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("error", echo.ErrBadRequest.Code, "error request", msgError)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	validate := validator.New()
	err = validate.Struct(checkemail)
	if err != nil {
		errors := helper.FormatError(err)
		msgError := echo.Map{
			"errors": errors,
		}
		response := helper.APIResponse("unprocassable", echo.ErrUnprocessableEntity.Code, "error", msgError)
		return c.JSON(echo.ErrUnprocessableEntity.Code, response)
	}
	isemailavailable, err := h.service.IsEmailAvailable(checkemail)
	if !isemailavailable || err != nil {
		msgerr := echo.Map{
			"errors": "email is not available",
		}
		response := helper.APIResponse("failed", echo.ErrNotFound.Code, "failed", msgerr)
		return c.JSON(echo.ErrNotFound.Code, response)
	}
	msg := echo.Map{
		"is available": isemailavailable,
	}
	response := helper.APIResponse("email is available", http.StatusOK, "success", msg)
	return c.JSON(http.StatusOK, response)
}
func (h *handler) UploadAvatar(c echo.Context) error {
	// user input avatar di passing ke struct input
	// struct input dipassing ke service
	// service dipasssing ke repository
	// repository -db
	file, err := c.FormFile("avatar")
	if err != nil {
		msg := echo.Map{
			"is uploaded": false,
		}
		response := helper.APIResponse("bad request", echo.ErrBadRequest.Code, "error", msg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	src, err := file.Open()
	if err != nil {
		msg := echo.Map{
			"is uploaded": false,
		}
		response := helper.APIResponse("bad request", echo.ErrBadRequest.Code, "error", msg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	defer src.Close()
	dir := filepath.Join("images", "avatar")
	os.MkdirAll(dir, 0755)
	filename := file.Filename
	ext := filepath.Ext(filename)
	deleteExt := strings.TrimSuffix(filename, ext)
	newFileName := deleteExt + ".png"
	dst, err := os.Create(filepath.Join(dir, newFileName))
	if err != nil {
		msg := echo.Map{
			"is uploaded": false,
		}
		response := helper.APIResponse("bad request", echo.ErrBadRequest.Code, "error", msg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		msg := echo.Map{
			"is uploaded": false,
		}
		response := helper.APIResponse("bad request", echo.ErrBadRequest.Code, "error", msg)
		return c.JSON(echo.ErrBadRequest.Code, response)
	}
	data := echo.Map{
		"is uploaded": true,
		"image url":   newFileName,
	}
	response := helper.APIResponse("success uploaded", http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, response)
}
