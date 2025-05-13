package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/achmad-dev/simple-ewallet/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserStockHoldRepository interface {
	GetByUserIDAndStockSymbol(ctx context.Context, userID string, stockSymbol string) (*domain.UserStockHolding, error)
	ListByUserID(ctx context.Context, userID string) ([]*domain.UserStockHolding, error)
	Upsert(ctx context.Context, holding *domain.UserStockHolding) (*domain.UserStockHolding, error)
}

type userStockHoldRepository struct {
	db *sqlx.DB
}

func NewUserStockHoldRepository(db *sqlx.DB) UserStockHoldRepository {
	return &userStockHoldRepository{
		db: db,
	}
}

const getUserByUserIDAndStockSymbolQuery = `
SELECT
    ush.id,
    ush.user_id,
    ush.stock_id,
    ush.quantity,
    ush.average_purchase_price,
    ush.created_at,
    ush.updated_at,
    s.symbol as stock_symbol,
    s.name as stock_name
FROM
    user_stock_holdings ush
JOIN
    stocks s ON ush.stock_id = s.id
WHERE
    ush.user_id = $1 AND s.symbol = $2;`

func (r *userStockHoldRepository) GetByUserIDAndStockSymbol(ctx context.Context, userID string, stockSymbol string) (*domain.UserStockHolding, error) {
	var holding domain.UserStockHolding
	err := r.db.GetContext(ctx, &holding, getUserByUserIDAndStockSymbolQuery, userID, stockSymbol)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user stock holding by user ID %s and stock symbol %s: %w", userID, stockSymbol, err)
	}
	return &holding, nil
}

const listHoldingsByUserIDQuery = `
SELECT
    ush.id,
    ush.user_id,
    ush.stock_id,
    ush.quantity,
    ush.average_purchase_price,
    ush.created_at,
    ush.updated_at,
    s.symbol as stock_symbol,
    s.name as stock_name
FROM
    user_stock_holdings ush
JOIN
    stocks s ON ush.stock_id = s.id
WHERE
    ush.user_id = $1
ORDER BY s.symbol ASC;`

func (r *userStockHoldRepository) ListByUserID(ctx context.Context, userID string) ([]*domain.UserStockHolding, error) {
	var holdings []*domain.UserStockHolding
	err := r.db.SelectContext(ctx, &holdings, listHoldingsByUserIDQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list user stock holdings for user ID %s: %w", userID, err)
	}
	if holdings == nil {
		return []*domain.UserStockHolding{}, nil
	}
	return holdings, nil
}

const upsertUserStockHoldingQuery = `
INSERT INTO user_stock_holdings (id, user_id, stock_id, quantity, average_purchase_price, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (user_id, stock_id) DO UPDATE SET
    quantity = EXCLUDED.quantity,
    average_purchase_price = EXCLUDED.average_purchase_price,
    updated_at = EXCLUDED.updated_at
RETURNING id, user_id, stock_id, quantity, average_purchase_price, created_at, updated_at;`

func (r *userStockHoldRepository) Upsert(ctx context.Context, holding *domain.UserStockHolding) (*domain.UserStockHolding, error) {
	if holding.ID == "" {
		holding.ID = uuid.NewString()
	}
	now := time.Now()
	if holding.CreatedAt.IsZero() {
		holding.CreatedAt = now
	}
	holding.UpdatedAt = now

	var upsertedHolding domain.UserStockHolding
	err := r.db.QueryRowxContext(
		ctx,
		upsertUserStockHoldingQuery,
		holding.ID,
		holding.UserID,
		holding.StockID,
		holding.Quantity,
		holding.AveragePurchasePrice,
		holding.CreatedAt,
		holding.UpdatedAt,
	).StructScan(&upsertedHolding)

	if err != nil {
		return nil, fmt.Errorf("failed to upsert user stock holding for user ID %s and stock ID %s: %w", holding.UserID, holding.StockID, err)
	}

	return &upsertedHolding, nil
}
