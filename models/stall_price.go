package models

import "time"

type StallPrice struct {
	ID      uint      `json:"id"`
	Price   float64   `json:"price"`
	Date    time.Time `json:"date"`
	StallID string    `json:"stall_id"`
}
