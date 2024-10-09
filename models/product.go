package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
	Likes       int       `json:"like"`
	Viewers     int       `json:"view"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
