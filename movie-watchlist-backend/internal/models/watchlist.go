package models

import "time"

type WatchStatus string

const (
	StatusWantToWatch WatchStatus = "Want to Watch"
	StatusWatching    WatchStatus = "Watching"
	StatusWatched     WatchStatus = "Watched"
	StatusSkipped     WatchStatus = "Skipped"
)

type Watchlist struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `gorm:"index;not null" json:"user_id"`
	MovieID   uint        `gorm:"index;not null" json:"movie_id"` // Matches TMDB ID
	Status    WatchStatus `gorm:"type:varchar(20);not null;default:'Want to Watch'" json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`

	// Optional relationships if doing joined queries
	User  *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Movie *Movie `gorm:"foreignKey:MovieID" json:"movie,omitempty"`
}
