package domain

import "time"

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/
type EWallet struct {
	ID        string     `json:"id" db:"id"`
	OwnerID   string     `json:"owner_id" db:"owner_id"`
	Balance   float64    `json:"balance" db:"balance"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
