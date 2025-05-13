package service

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/

import (
	"context"
	"errors"

	"github.com/achmad-dev/simple-ewallet/internal/domain"
	"github.com/achmad-dev/simple-ewallet/internal/dto"
	"github.com/achmad-dev/simple-ewallet/internal/pkg"
	"github.com/achmad-dev/simple-ewallet/internal/repository"
	"github.com/sirupsen/logrus"
)

// service interface for user
type UserService interface {
	SignIn(ctx context.Context, username, password string) (*dto.AuthResponseDto, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	Signup(ctx context.Context, userSignup dto.UserSignupDto) (*dto.AuthResponseDto, error)
}

// user service implementation
type userServiceImpl struct {
	userRepo    repository.UserRepo
	ewalletRepo repository.EWalletRepository
	bcryptUtil  pkg.BcryptUtil
	secret      string
	log         *logrus.Logger
}

// Signup implements UserService.
func (u *userServiceImpl) Signup(ctx context.Context, userSignup dto.UserSignupDto) (*dto.AuthResponseDto, error) {
	_, err := u.userRepo.GetUserByName(ctx, userSignup.Username)
	if err == nil {
		u.log.Error("user already exists")
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := u.bcryptUtil.HashPassword(userSignup.Password)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("something went wrong")
	}

	newUser := &domain.User{
		Username: userSignup.Username,
		Email:    userSignup.Email,
		Password: hashedPassword,
	}

	id, err := u.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("something went wrong")
	}

	// create ewallet for user
	err = u.ewalletRepo.AddEWallet(id)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("something went wrong")
	}

	token, err := pkg.GenerateToken(newUser.ID, u.secret)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("something went wrong")
	}

	response := &dto.AuthResponseDto{
		Token: token,
	}

	return response, nil
}

// GetUserByID implements UserService.
func (u *userServiceImpl) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("user not found")
	}
	return user, nil
}

// SignIn implements UserService.
func (u *userServiceImpl) SignIn(ctx context.Context, username string, password string) (*dto.AuthResponseDto, error) {
	user, err := u.userRepo.GetUserByName(ctx, username)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("user not found")
	}
	if !u.bcryptUtil.CheckPasswordHash(password, user.Password) {
		u.log.Error("invalid password")
		return nil, errors.New("invalid password")
	}
	token, err := pkg.GenerateToken(user.ID, u.secret)
	if err != nil {
		u.log.Error(err)
		return nil, errors.New("something went wrong")
	}

	response := &dto.AuthResponseDto{
		Token: token,
	}

	return response, nil
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepo, ewalletRepo repository.EWalletRepository, bcryptUtil pkg.BcryptUtil, secret string, log *logrus.Logger) UserService {
	return &userServiceImpl{userRepo: userRepo, ewalletRepo: ewalletRepo, bcryptUtil: bcryptUtil, secret: secret, log: log}
}
