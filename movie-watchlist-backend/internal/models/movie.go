package models

import (
	"time"
)

// Movie represents cached movie metadata from TMDB/OMDB locally
type Movie struct {
	ID          uint      `gorm:"primaryKey" json:"id"` // TMDB ID or custom ID
	Title       string    `gorm:"index;not null" json:"title"`
	Overview    string    `json:"overview"`
	ReleaseDate string    `json:"release_date"`
	PosterPath  string    `json:"poster_path"`
	GenreIDs    string    `json:"genre_ids"` // Stored as comma-separated or JSON
	CachedAt    time.Time `json:"cached_at"`
}
