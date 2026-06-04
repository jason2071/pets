package service

import (
	"errors"

	"github.com/jason2071/pets/internal/auth"
	"github.com/jason2071/pets/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrAccountNotFound    = errors.New("account not found")
	ErrEmailExists        = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type AccountService struct {
	repo   domain.AccountRepository
	tokens *auth.TokenManager
}

func NewAccountService(repo domain.AccountRepository, tokens *auth.TokenManager) *AccountService {
	return &AccountService{repo: repo, tokens: tokens}
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

// Login authenticates by email + plaintext password and returns a signed JWT.
// Both "no such email" and "wrong password" return ErrInvalidCredentials to
// avoid user enumeration.
func (s *AccountService) Login(email, plain string) (string, error) {
	found, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if !s.VerifyPassword(found, plain) {
		return "", ErrInvalidCredentials
	}

	return s.tokens.Generate(found.ID, found.Name, found.Email, found.Role)
}
