package models

import "time"

type Rating struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	MovieID   uint      `gorm:"index;not null" json:"movie_id"`
	Score     int       `gorm:"not null;check:score >= 1 AND score <= 5" json:"score"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	User  *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Movie *Movie `gorm:"foreignKey:MovieID" json:"movie,omitempty"`
}
