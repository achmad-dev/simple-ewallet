package repository

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/

import (
	"context"

	"github.com/achmad-dev/simple-ewallet/internal/domain"
	"github.com/jmoiron/sqlx"
)

// base interface for user repository
type UserRepo interface {
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	GetUserByName(ctx context.Context, username string) (*domain.User, error)
	CreateUser(ctx context.Context, newUser *domain.User) error
}

// user repository implementation
type userRepositoryImpl struct {
	sqlx *sqlx.DB
}

// CreateUser implements UserRepo.
func (u *userRepositoryImpl) CreateUser(ctx context.Context, newUser *domain.User) error {
	query := `INSERT INTO users (username, password, email, created_at, updated_at) VALUES ($1, $2, $3)`
	_, err := u.sqlx.ExecContext(ctx, query,
		newUser.Username,
		newUser.Password,
		newUser.Email,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByName implements UserRepo.
func (u *userRepositoryImpl) GetUserByName(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, username, password, role, company_name, created_at, updated_at FROM users WHERE username = $1`
	err := u.sqlx.GetContext(ctx, user, query, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID implements UserRepo.
func (u *userRepositoryImpl) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, username, password, role, company_name, created_at, updated_at FROM users WHERE id = $1`
	err := u.sqlx.GetContext(ctx, user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// NewUserRepository creates a new user repository
func NewUserRepository(sqlx *sqlx.DB) UserRepo {
	return &userRepositoryImpl{sqlx: sqlx}
}
