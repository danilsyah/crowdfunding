package main

import (
	"crowdfunding/auth"
	"crowdfunding/campaign"
	"crowdfunding/handler"
	"crowdfunding/helper"
	"crowdfunding/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ==== analisa arsitektur ====
// input dari user
// handler, mapping input dari user -> struct input
// service : melakukan mapping dari struct input ke struct user
// repository : simpan ke database
// database : simpan ke database
func main() {
	// gin.SetMode(gin.ReleaseMode)

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Get port from env or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not specified
	}
	// Build database connection string from env variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	fmt.Println("Connected to database!")

	// repository
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	// service
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewJWTService()

	// handler
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	// endpoint users
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// endpoint campaigns
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaignDetail)

	fmt.Printf("Server starting on port :%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("failed to start server:", err)
	}
}

// ==== analisa auth middleware ====
// ambil nilai header Authorization: Bearer tokentokentokentoken
// dari header Authorization, ambil nilai tokennya
// kita validasi token
// kita ambil user_id dari token
// ambil user dari database berdasarkan user_id lewat service
// kita set context isinya user
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}
