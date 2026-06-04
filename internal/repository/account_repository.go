package repository

import (
	"github.com/jason2071/pets/internal/domain"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) domain.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(a *domain.Account) error {
	return nil
}

func (r *accountRepository) FindById(id uint) (*domain.Account, error) {
	return nil, nil
}

func (r *accountRepository) FindByEmail(email string) (*domain.Account, error) {
	return nil, nil
}

func (r *accountRepository) Update(a *domain.Account) error {
	return nil
}

func (r *accountRepository) Delete(id uint) error {
	return nil
}
