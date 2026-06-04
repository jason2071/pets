package service

import (
	"errors"

	"github.com/jason2071/pets/internal/domain"
)

var ErrAccountNotFound = errors.New("account not found")

type AccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Create(acc *domain.Account) error {
	return s.repo.Create(acc)
}
