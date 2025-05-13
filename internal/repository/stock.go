package repository

import (
	"github.com/achmad-dev/simple-ewallet/internal/domain"
	"github.com/jmoiron/sqlx"
)

type StockRepository interface {
	UpdateStockPriceByName(name string, price float64) error
	GetAllStockPrice() ([]*domain.Stock, error)
}

type stockRepository struct {
	sqlx *sqlx.DB
}

// GetAllStockPrice implements StockRepository.
func (s *stockRepository) GetAllStockPrice() ([]*domain.Stock, error) {
	panic("unimplemented")
}

// UpdateStockPriceByName implements StockRepository.
func (s *stockRepository) UpdateStockPriceByName(name string, price float64) error {
	panic("unimplemented")
}

func NewStockRepository(sqlx *sqlx.DB) StockRepository {
	return &stockRepository{
		sqlx: sqlx,
	}
}
