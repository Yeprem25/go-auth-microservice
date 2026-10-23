package main

import (
	"cmp"
	"log"
	"net/http"
	"os"

	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/database"
	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	port := cmp.Or(os.Getenv("PORT"), "8080")
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
