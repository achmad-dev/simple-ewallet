package domain

import "time" // Added import

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/
type UserStockHolding struct {
	ID                   string    `json:"id" db:"id"`
	UserID               string    `json:"user_id" db:"user_id"`
	StockID              string    `json:"stock_id" db:"stock_id"`
	Quantity             float64   `json:"quantity" db:"quantity"`
	AveragePurchasePrice float64   `json:"average_purchase_price" db:"average_purchase_price"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`

	StockSymbol string `json:"stock_symbol,omitempty" db:"stock_symbol"`
	StockName   string `json:"stock_name,omitempty" db:"stock_name"`
}
