package domain

import "time"

// Pet is the core domain entity.
type Pet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OwnerID   *uint     `gorm:"index" json:"owner_id,omitempty"`
	Name      string    `gorm:"not null" json:"name"`
	Species   string    `gorm:"not null" json:"species"`
	Breed     string    `json:"breed,omitempty"`
	Age       int       `json:"age,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PetRepository defines persistence operations for pets.
type PetRepository interface {
	Create(p *Pet) error
	FindAll() ([]Pet, error)
	FindByID(id uint) (*Pet, error)
	Update(p *Pet) error
	Delete(id uint) error
}
