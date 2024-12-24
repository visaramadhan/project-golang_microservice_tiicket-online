package models

import "time"

// Category model
type Category struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	IconURL     string     `json:"icon_url"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
}
