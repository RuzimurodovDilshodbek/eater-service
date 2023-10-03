package models

import "time"

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
type Address struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	EaterID   string    `json:"eater_id"` // reference to eater
	Name      string    `json:"name"`
	Location  *Location `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
