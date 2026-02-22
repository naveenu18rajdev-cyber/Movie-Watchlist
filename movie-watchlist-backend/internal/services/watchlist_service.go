package services

import (
	"movie-watchlist-backend/internal/models"
	"movie-watchlist-backend/internal/repositories"
)

type WatchlistService struct {
	watchlistRepo *repositories.WatchlistRepository
}

func NewWatchlistService(repo *repositories.WatchlistRepository) *WatchlistService {
	return &WatchlistService{
		watchlistRepo: repo,
	}
}

func (s *WatchlistService) AddToWatchlist(userID uint, movieID uint, status string) (*models.Watchlist, error) {
	parsedStatus := models.StatusWantToWatch
	if status != "" {
		parsedStatus = models.WatchStatus(status)
	}

	item := &models.Watchlist{
		UserID:  userID,
		MovieID: movieID,
		Status:  parsedStatus,
	}

	err := s.watchlistRepo.Create(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *WatchlistService) UpdateStatus(id uint, userID uint, status string) error {
	return s.watchlistRepo.UpdateStatus(id, userID, models.WatchStatus(status))
}

func (s *WatchlistService) RemoveFromWatchlist(id uint, userID uint) error {
	return s.watchlistRepo.Delete(id, userID)
}

func (s *WatchlistService) GetWatchlist(userID uint) ([]*models.Watchlist, error) {
	return s.watchlistRepo.GetUserWatchlist(userID)
}
