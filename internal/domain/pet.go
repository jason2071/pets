package domain

import "time"

// Pet is the core domain entity.
type Pet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OwnerID   *uint     `gorm:"index" json:"owner_id,omitempty"`
	Owner     *Account  `gorm:"constraint:OnDelete:SET NULL" json:"owner,omitempty"`
	Name      string    `gorm:"not null" json:"name"`
	Species   string    `gorm:"not null" json:"species"`
	Breed     string    `json:"breed,omitempty"`
	Age       int       `json:"age,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PetRepository defines persistence operations for pets. Reads, updates, and
// deletes are scoped by owner so accounts only touch their own pets.
type PetRepository interface {
	Create(p *Pet) error
	FindAllByOwner(ownerID uint) ([]Pet, error)
	FindByIDAndOwner(id, ownerID uint) (*Pet, error)
	Update(p *Pet) error
	DeleteByOwner(id, ownerID uint) error
}
