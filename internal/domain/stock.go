package domain

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/

type Stock struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
