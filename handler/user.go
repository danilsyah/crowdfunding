package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

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

func generateRandomString(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// user mengupload avatar lewat http multipart form-data
	// handler menerima file dari user
	// handler memvalidasi file yang diupload
	// handler memanggil service untuk menyimpan file ke storage
	// handler memanggil service untuk update user dengan avatar yang baru
	// handler mengembalikan response ke user

	file, err := c.FormFile("avatar")
	if err != nil {
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get file extension and validate
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	if !allowedExts[ext] {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Invalid file format. Only JPG, JPEG, PNG and GIF are allowed", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// validation file image upload size = 1 MB = 1024 * 1024 bytes
	if file.Size > 1024*1024 {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Avatar image is too large", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1 // ambil dari token jwt
	// Get current avatar path
	currentUser, err := h.userService.GetUserByID(userID)

	// generate random file name
	randomFileName := generateRandomString(10)
	newFilename := fmt.Sprintf("%d_%s%d%s", userID, randomFileName, time.Now().Unix(), ext)

	// Ensure images directory exists
	if err := os.MkdirAll("images", 0755); err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to create upload directory", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Delete old avatar file if exists
	if err == nil && currentUser.AvatarFileName != "" {
		oldAvatarPath := filepath.Join(currentUser.AvatarFileName)
		fmt.Println("Deleting old avatar:", oldAvatarPath)
		if _, err := os.Stat(oldAvatarPath); err == nil {
			os.Remove(oldAvatarPath)
		}
	}

	// Save new Avatar file
	path := filepath.Join("images", newFilename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
