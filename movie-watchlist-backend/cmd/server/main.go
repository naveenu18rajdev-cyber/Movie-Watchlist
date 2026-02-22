package main

import (
	"log"

	"movie-watchlist-backend/internal/cache"
	"movie-watchlist-backend/internal/config"
	"movie-watchlist-backend/internal/database"
	"movie-watchlist-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Configuration
	cfg := config.LoadConfig()

	// Initialize Database and Run Migrations
	database.ConnectDB(cfg)
	database.RunMigrations()

	// Initialize Redis Cache
	cache.ConnectRedis(cfg)

	// Initialize Gin router
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
			"message": "Movie Watchlist Platform API is running",
		})
	})

	// Initialize and bind domain handlers to routing group
	routes.SetupRoutes(r, cfg)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
