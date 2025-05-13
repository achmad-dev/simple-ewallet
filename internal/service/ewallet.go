package service

import (
	"errors"

	"github.com/achmad-dev/simple-ewallet/internal/domain"
	"github.com/achmad-dev/simple-ewallet/internal/repository"
	"github.com/sirupsen/logrus"
)

type EWalletService interface {
	AddBalance(userID string, amount float64) error
	SubtractBalance(userID string, amount float64) error
	GetWallet(userID string) (*domain.EWallet, error)
}

type eWalletService struct {
	ewalletRepository repository.EWalletRepository
	log               *logrus.Logger
}

// AddBalance implements EWalletService.
func (e *eWalletService) AddBalance(userID string, amount float64) error {
	ewallet, err := e.ewalletRepository.GetEWallet(userID)
	if err != nil {
		return err
	}

	err = e.ewalletRepository.UpdateEWallet(userID, ewallet, amount)
	if err != nil {
		e.log.Error(err)
		return err
	}
	return nil
}

// GetWallet implements EWalletService.
func (e *eWalletService) GetWallet(userID string) (*domain.EWallet, error) {
	ewallet, err := e.ewalletRepository.GetEWallet(userID)
	if err != nil {
		e.log.Error(err)
		return nil, err
	}
	return ewallet, nil
}

// SubtractBalance implements EWalletService.
func (e *eWalletService) SubtractBalance(userID string, amount float64) error {
	ewallet, err := e.ewalletRepository.GetEWallet(userID)
	if err != nil {
		return err
	}
	if ewallet.Balance < amount {
		return errors.New("insufficient balance")
	}
	err = e.ewalletRepository.UpdateEWallet(userID, ewallet, -amount)
	if err != nil {
		e.log.Error(err)
		return err
	}
	return nil
}

func NewEWalletService(ewalletRepository repository.EWalletRepository, log *logrus.Logger) EWalletService {
	return &eWalletService{
		ewalletRepository: ewalletRepository,
		log:               log,
	}
}
