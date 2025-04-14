package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// RegisterUser adalah handler untuk register user
// @Summary Register User
func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap inputan dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter ke service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to register account", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Failed to register account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentokentoken")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan input email dan password
	// input ditangkap handler
	// mapping dari input user ke input struct
	// input struct passing service
	// di service mencari dengan bantuan repository user dengan email x
	// mencocokan password

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentokentoken")
	response := helper.APIResponse("Successfully logged in", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// ada input email dari user
	// input email di-mapping ke struct
	// struct input di-passing ke service
	// service akan manggil repository - email sudah ada atau belum
	// repository akan mencari email di database

	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to check email availability", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"error": "Server Error"}
		response := helper.APIResponse("Failed to check email availability", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}

	metaMessage := "Email has been available"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
