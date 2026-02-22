package database

import (
	"log"

	"movie-watchlist-backend/internal/models"
)

// RunMigrations automigrates all defined GORM models
func RunMigrations() {
	if DB == nil {
		log.Fatal("Database connection not initialized before running migrations")
	}

	err := DB.AutoMigrate(
		&models.User{},
		&models.Movie{},
		&models.Watchlist{},
		&models.Rating{},
	)

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully")
}
