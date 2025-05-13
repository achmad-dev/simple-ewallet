package domain

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/
type EWallet struct {
	ID         string  `json:"id"`
	Balance    float64 `json:"balance"`
	OwnerID    string  `json:"owner_id"`
	OwnerName  string  `json:"owner_name"`
	OwnerEmail string  `json:"owner_email"`
}
