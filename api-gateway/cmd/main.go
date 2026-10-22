package main

import (
	"cmp"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
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
