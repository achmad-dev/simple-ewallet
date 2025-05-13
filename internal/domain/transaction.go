package domain

type EWalletTransaction struct {
	ID                 int64    `json:"id"`
	WalletID           int64    `json:"wallet_id"`
	TransactionType    string   `json:"transaction_type"`     // "add" or "withdraw" or "buy_stock"
	RelatedStockID     *int64   `json:"related_stock_id"`     // nil if not related to a stock
	RelatedStockName   *string  `json:"related_stock_name"`   // Name of the stock if related
	PriceAtTransaction *float64 `json:"price_at_transaction"` // Price of the stock at the time of transaction
	Amount             float64  `json:"amount"`
	Description        string   `json:"description"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
}
