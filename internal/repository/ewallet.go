package repository

import (
	"fmt"

	"github.com/achmad-dev/simple-ewallet/internal/domain"
	"github.com/jmoiron/sqlx"
)

type EWalletRepository interface {
	// AddEWallet adds a new eWallet to the database
	AddEWallet(userID string) error
	// GetEWallet retrieves the eWallet for a given userID
	GetEWallet(userID string) (*domain.EWallet, error)
	// UpdateEWallet updates the eWallet for a given userID
	UpdateEWallet(userID string, eWallet *domain.EWallet, amount float64) error
}

type eWalletRepository struct {
	sqlx *sqlx.DB
}

// AddEWallet implements EWalletRepository.
func (e *eWalletRepository) AddEWallet(userID string) error {
	exec := `
		INSERT INTO ewallets (owner_id, balance, created_at, updated_at)
		VALUES ($1, 0, NOW(), NOW())
		ON CONFLICT (owner_id) DO NOTHING
	`
	_, err := e.sqlx.Exec(exec, userID)
	if err != nil {
		return fmt.Errorf("failed to add eWallet: %w", err)
	}
	return nil
}

// GetEWallet implements EWalletRepository.
func (e *eWalletRepository) GetEWallet(userID string) (*domain.EWallet, error) {
	exec := `
		SELECT id, owner_id, balance
		FROM ewallets
		WHERE owner_id = $1
	`
	ewallet := &domain.EWallet{}
	err := e.sqlx.QueryRow(exec, userID).Scan(
		&ewallet.ID,
		&ewallet.OwnerID,
		&ewallet.Balance,
	)
	if err != nil {
		return nil, err
	}
	return ewallet, nil
}

// UpdateEWallet implements EWalletRepository.
func (e *eWalletRepository) UpdateEWallet(userID string, eWallet *domain.EWallet, amount float64) error {
	tx, err := e.sqlx.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	var currentBalance float64
	err = tx.Get(&currentBalance, "SELECT balance FROM ewallets WHERE owner_id = $1 FOR UPDATE", userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if amount < 0 && currentBalance <= 0 {
		tx.Rollback()
		return fmt.Errorf("insufficient balance to withdraw")
	}

	newBalance := currentBalance + amount
	if newBalance < 0 {
		tx.Rollback()
		return fmt.Errorf("insufficient balance to withdraw")
	}

	_, err = tx.Exec(
		"UPDATE ewallets SET balance = $1, updated_at = NOW() WHERE owner_id = $2",
		newBalance, userID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	eWallet.Balance = newBalance
	return nil
}

func NewEWalletRepository(sqlx *sqlx.DB) EWalletRepository {
	return &eWalletRepository{
		sqlx: sqlx,
	}
}
