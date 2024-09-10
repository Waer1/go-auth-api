package models

import (
	"time"
)

// BaseEntity defines common fields for all entities.
type BaseEntity struct {
	ID        uint      `gorm:"primaryKey" json:"id"` // Primary key ID
	CreatedAt time.Time `json:"-"`                    // Exclude CreatedAt from JSON
	UpdatedAt time.Time `json:"-"`                    // Exclude UpdatedAt from JSON
}
