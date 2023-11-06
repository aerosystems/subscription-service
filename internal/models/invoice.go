package models

import "time"

type Invoice struct {
	Id        int       `json:"id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
