package interfaces

import (
	"time"
)

type WeddingDTO struct {
	ID         string    `json:"id"`
	Couple     Couple    `json:"couple"`
	Date       time.Time `json:"date"`
	Location   string    `json:"location"`
	GuestCount int       `json:"guest_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
