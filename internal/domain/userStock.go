package domain

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/
type UserStockHolding struct {
	ID         string  `json:"id"`
	StockID    string  `json:"stock_id"`
	StockName  string  `json:"stock_name"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
	TotalValue float64 `json:"total_value"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
