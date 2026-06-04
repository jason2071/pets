package service

import (
	"errors"

	"github.com/jason2071/pets/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

var ErrAccountNotFound = errors.New("account not found")

type AccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Create(acc *domain.Account) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	acc.Password = string(hash)

	return s.repo.Create(acc)
}

// VerifyPassword checks a plaintext password against the stored bcrypt hash.
func (s *AccountService) VerifyPassword(acc *domain.Account, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(plain)) == nil
}
