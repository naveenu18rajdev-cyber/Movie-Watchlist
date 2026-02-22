package repositories

import (
	"movie-watchlist-backend/internal/database"
	"movie-watchlist-backend/internal/models"
)

type MovieRepository struct{}

func NewMovieRepository() *MovieRepository {
	return &MovieRepository{}
}

func (r *MovieRepository) Save(movie *models.Movie) error {
	return database.DB.Save(movie).Error
}

func (r *MovieRepository) FindByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	err := database.DB.First(&movie, id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}
