package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Yusufdot101/eventhive/internal/auth"
	"github.com/Yusufdot101/eventhive/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.SetupVars()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     config.AllowedOrigins,
		AllowMethods:     config.AllowedMethods,
		AllowHeaders:     config.AllowedHeaders,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "active",
		})
	})

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
