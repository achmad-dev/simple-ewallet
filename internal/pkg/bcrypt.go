package pkg

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptUtil interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type bcryptUtilImpl struct {
	cost int
}

func NewBcryptUtil(cost int) BcryptUtil {
	return &bcryptUtilImpl{cost: cost}
}

func (b *bcryptUtilImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	return string(bytes), err
}

func (b *bcryptUtilImpl) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
