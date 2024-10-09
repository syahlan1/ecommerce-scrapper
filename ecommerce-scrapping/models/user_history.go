package models

import "time"

type UserHistory struct {
	ID              uint      `gorm:"primaryKey"`
	UserID          int       `json:"user_id"`
	ProductID       int       `json:"product_id,omitempty"`
	ProductName     string    `json:"product_name,omitempty"`
	InteractionType string    `json:"interaction_type"` // 'search', 'view'
	Price           int       `json:"price"`
	Category	string	`json:"category"`
	SearchQuery     string    `json:"search_query,omitempty"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
}
