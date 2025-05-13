package domain

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/
type EWalletTransaction struct {
	ID                 string   `json:"id"`
	WalletID           string   `json:"wallet_id"`
	TransactionType    string   `json:"transaction_type"`     // "add" or "withdraw" or "buy_stock"
	RelatedStockID     *string  `json:"related_stock_id"`     // nil if not related to a stock
	RelatedStockName   *string  `json:"related_stock_name"`   // Name of the stock if related
	PriceAtTransaction *float64 `json:"price_at_transaction"` // Price of the stock at the time of transaction
	QuantityTransacted *float64 `json:"quantity_transacted"`  // Quantity of stock transacted, nil if not related to a stock
	Amount             float64  `json:"amount"`
	Description        string   `json:"description"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
}
