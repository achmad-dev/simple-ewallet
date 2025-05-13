package domain

type EWallet struct {
	ID         int64   `json:"id"`
	Balance    float64 `json:"balance"`
	OwnerID    int64   `json:"owner_id"`
	OwnerName  string  `json:"owner_name"`
	OwnerEmail string  `json:"owner_email"`
}
