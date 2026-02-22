package repositories

import (
	"movie-watchlist-backend/internal/database"
	"movie-watchlist-backend/internal/models"
)

type RatingRepository struct{}

func NewRatingRepository() *RatingRepository {
	return &RatingRepository{}
}

func (r *RatingRepository) Create(rating *models.Rating) error {
	return database.DB.Create(rating).Error
}

func (r *RatingRepository) GetUserRatings(userID uint) ([]*models.Rating, error) {
	var ratings []*models.Rating
	err := database.DB.Where("user_id = ?", userID).Find(&ratings).Error
	return ratings, err
}
