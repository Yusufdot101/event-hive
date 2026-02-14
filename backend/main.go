package main

import (
	"log"
	"net/http"

	"github.com/Yusufdot101/eventhive/internal/auth"
	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "active",
		})
	})

	config.SetupVars()
	DB, err := config.OpenDB(config.DSN)
	if err != nil {
		log.Fatalf("error opening DB: %v", err)
	}
	group := r.Group("/auth")

	auth.RegisterRoutes(DB, group)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}
