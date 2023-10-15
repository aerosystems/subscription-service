package models

import "time"

type Invoice struct {
	Id        uint      `json:"id"`
	Amount    uint      `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
