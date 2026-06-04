package domain

import "time"

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email"`
	Password  string    `gorm:"column:password_hash;not null" json:"-"`
	Name      string    `gorm:"column:full_name;not null" json:"name"`
	Role      string    `gorm:"not null;default:owner" json:"role"`
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
