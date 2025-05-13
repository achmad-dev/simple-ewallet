package repository

import (
	"context"

	"github.com/achmad-dev/simple-ewallet/internal/domain"
	"github.com/jmoiron/sqlx"
)

type EWalletTransactionRepository interface {
	AddTransaction(ctx context.Context, transaction domain.EWalletTransaction) error
}

type eWalletTransactionRepository struct {
	sqlx *sqlx.DB
}

// AddTransaction implements EWalletTransactionRepository.
func (e *eWalletTransactionRepository) AddTransaction(ctx context.Context, transaction domain.EWalletTransaction) error {
	query := `
		INSERT INTO ewallet_transactions (
			id, wallet_id, transaction_type, amount, description,
			related_stock_id, price_at_transaction, quantity_transacted, created_at, updated_at
		) VALUES (
			COALESCE(?, gen_random_uuid()), ?, ?, ?, ?,
			?, ?, ?, COALESCE(?, now()), COALESCE(?, now())
		)
	`

	_, err := e.sqlx.ExecContext(
		ctx,
		query,
		transaction.ID,
		transaction.WalletID,
		transaction.TransactionType,
		transaction.Amount,
		transaction.Description,
		transaction.RelatedStockID,
		transaction.PriceAtTransaction,
		transaction.QuantityTransacted,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)
	return err
}

func NewEWalletTransactionRepository(sqlx *sqlx.DB) EWalletTransactionRepository {
	return &eWalletTransactionRepository{
		sqlx: sqlx,
	}
}
