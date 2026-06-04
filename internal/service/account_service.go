package service

import (
	"errors"

	"github.com/jason2071/pets/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrEmailExists     = errors.New("email already registered")
)

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

	if err := s.repo.Create(acc); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrEmailExists
		}
		return err
	}
	return nil
}

// VerifyPassword checks a plaintext password against the stored bcrypt hash.
func (s *AccountService) VerifyPassword(acc *domain.Account, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(plain)) == nil
}

// Login authenticates by email + plaintext password. acc carries the input
// (Email + plaintext Password); on success it is replaced with the stored
// account (id, hashed password, etc.).
func (s *AccountService) Login(acc *domain.Account) error {
	plain := acc.Password

	found, err := s.repo.FindByEmail(acc.Email)
	if err != nil {
		// Unify "no such email" and "wrong password" to avoid user enumeration.
		return ErrAccountNotFound
	}

	if !s.VerifyPassword(found, plain) {
		return ErrAccountNotFound
	}

	*acc = *found
	return nil
}
