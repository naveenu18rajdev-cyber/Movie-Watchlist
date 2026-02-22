package services

import (
	"encoding/json"
	"fmt"
	"time"

	"movie-watchlist-backend/internal/cache"
	"movie-watchlist-backend/internal/config"
	"movie-watchlist-backend/internal/external"
	"movie-watchlist-backend/internal/repositories"

	"github.com/redis/go-redis/v9"
)

type MovieService struct {
	tmdbClient *external.TMDBClient
	cache      *redis.Client
	config     *config.Config
	// To optionally store frequently accessed movies in DB to avoid rate limits
	movieRepo *repositories.MovieRepository 
}

func NewMovieService(tmdb *external.TMDBClient, rdb *redis.Client, cfg *config.Config, repo *repositories.MovieRepository) *MovieService {
	return &MovieService{
		tmdbClient: tmdb,
		cache:      rdb,
		config:     cfg,
		movieRepo:  repo,
	}
}

// SearchMovies checks cache, if miss then calls TMDB
func (s *MovieService) SearchMovies(query string) (map[string]interface{}, error) {
	cacheKey := fmt.Sprintf("movies:search:%s", query)
	
	// 1. Check Cache
	if s.cache != nil {
		if val, err := s.cache.Get(cache.Ctx, cacheKey).Result(); err == nil {
			var result map[string]interface{}
			if json.Unmarshal([]byte(val), &result) == nil {
				return result, nil // Cache Hit
			}
		}
	}

	// 2. Fetch from TMDB
	result, err := s.tmdbClient.SearchMovies(query)
	if err != nil {
		return nil, err
	}

	// 3. Store in Cache (TTL 1 Hour)
	if s.cache != nil {
		if bytes, err := json.Marshal(result); err == nil {
			s.cache.Set(cache.Ctx, cacheKey, string(bytes), 1*time.Hour)
		}
	}

	return result, nil
}

// GetMovieDetails checks cache, if miss calls TMDB 
func (s *MovieService) GetMovieDetails(movieID int) (map[string]interface{}, error) {
	cacheKey := fmt.Sprintf("movies:details:%d", movieID)
	
	if s.cache != nil {
		if val, err := s.cache.Get(cache.Ctx, cacheKey).Result(); err == nil {
			var result map[string]interface{}
			if json.Unmarshal([]byte(val), &result) == nil {
				return result, nil
			}
		}
	}

	result, err := s.tmdbClient.GetMovieDetails(movieID)
	if err != nil {
		return nil, err
	}

	// Store for 24 hours
	if s.cache != nil {
		if bytes, err := json.Marshal(result); err == nil {
			s.cache.Set(cache.Ctx, cacheKey, string(bytes), 24*time.Hour)
		}
	}

	return result, nil
}
