package repositories

import (
	"movie-watchlist-backend/internal/database"
	"movie-watchlist-backend/internal/models"
)

type WatchlistRepository struct{}

func NewWatchlistRepository() *WatchlistRepository {
	return &WatchlistRepository{}
}

func (r *WatchlistRepository) Create(item *models.Watchlist) error {
	return database.DB.Create(item).Error
}

func (r *WatchlistRepository) UpdateStatus(id uint, userID uint, status models.WatchStatus) error {
	return database.DB.Model(&models.Watchlist{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("status", status).Error
}

func (r *WatchlistRepository) Delete(id uint, userID uint) error {
	return database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Watchlist{}).Error
}

func (r *WatchlistRepository) GetUserWatchlist(userID uint) ([]*models.Watchlist, error) {
	var list []*models.Watchlist
	err := database.DB.Preload("Movie").Where("user_id = ?", userID).Find(&list).Error
	return list, err
}
