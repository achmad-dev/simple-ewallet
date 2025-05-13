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
	CreateUser(ctx context.Context, newUser *domain.User) (string, error)
}

// user repository implementation
type userRepositoryImpl struct {
	sqlx *sqlx.DB
}

// CreateUser implements UserRepo.
func (u *userRepositoryImpl) CreateUser(ctx context.Context, newUser *domain.User) (string, error) {
	query := `INSERT INTO users (username, password, email, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id` // Changed 'name' to 'username'
	var id string
	err := u.sqlx.QueryRowContext(ctx, query,
		newUser.Username,
		newUser.Password,
		newUser.Email,
	).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetUserByName implements UserRepo.
func (u *userRepositoryImpl) GetUserByName(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE username = $1` // Changed 'name' to 'username' and selected all fields
	err := u.sqlx.GetContext(ctx, user, query, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID implements UserRepo.
func (u *userRepositoryImpl) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = $1` // Selected all fields
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
