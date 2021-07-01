package handler

import (
	"blog/helper"
	"blog/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewHandlerUser(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMesage := gin.H{"error": errors}

		response := helper.ApiResponse("register failed", http.StatusUnprocessableEntity, "error", errorMesage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUserInput(input)
	if err != nil {
		response := helper.ApiResponse("register failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := user.FormatUser(newUser)
	response := helper.ApiResponse("register succes", http.StatusOK, "succes", formater)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UserLogin(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMesage := gin.H{"error": errors}

		response := helper.ApiResponse("login failed", http.StatusUnprocessableEntity, "error", errorMesage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedUser, err := h.userService.Login(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMesage := gin.H{"error": errors}

		response := helper.ApiResponse("login failed", http.StatusUnprocessableEntity, "error", errorMesage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formater := user.FormatUser(loggedUser)
	response := helper.ApiResponse("login sukses", http.StatusOK, "success", formater)
	c.JSON(http.StatusOK, response)
}
