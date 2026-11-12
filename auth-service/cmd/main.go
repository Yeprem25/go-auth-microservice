package main

import (
	"cmp"
	"log"
	"net/http"
	"os"

	"github.com/Yeprem25/go-auth-microservice/auth-service/internal/database"
	"github.com/Yeprem25/go-auth-microservice/auth-service/internal/handler"
	"github.com/Yeprem25/go-auth-microservice/auth-service/internal/middleware"
	"github.com/Yeprem25/go-auth-microservice/auth-service/internal/repository"
	"github.com/Yeprem25/go-auth-microservice/auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPostgresUserRepository(db)
	tokenRepo := repository.NewRefreshTokenRepository(db)
	authService := service.NewAuthService(repo, tokenRepo, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/refresh", authHandler.Refresh)
	r.POST("/logout", authHandler.Logout)

	protected := r.Group("/")
	protected.Use(middleware.Auth(authService))
	protected.GET("/me", authHandler.Me)

	port := cmp.Or(os.Getenv("PORT"), "8080")
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
