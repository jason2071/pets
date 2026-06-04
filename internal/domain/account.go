package domain

import "time"

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountRepository interface {
	Create(a *Account) error
	FindById(id uint) (*Account, error)
	FindByEmail(email string) (*Account, error)
	Update(a *Account) error
	Delete(id uint) error
}
