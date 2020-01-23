package game

import "time"

type Game struct {
	ID        int64
	Name      string
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
