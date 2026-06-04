package repository

import (
	"github.com/jason2071/pets/internal/domain"
	"gorm.io/gorm"
)

type petRepository struct {
	db *gorm.DB
}

// NewPetRepository returns a GORM-backed PetRepository.
func NewPetRepository(db *gorm.DB) domain.PetRepository {
	return &petRepository{db: db}
}

func (r *petRepository) Create(p *domain.Pet) error {
	return r.db.Create(p).Error
}

func (r *petRepository) FindAll() ([]domain.Pet, error) {
	var pets []domain.Pet
	err := r.db.Order("id").Find(&pets).Error
	return pets, err
}

func (r *petRepository) FindByID(id uint) (*domain.Pet, error) {
	var p domain.Pet
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *petRepository) Update(p *domain.Pet) error {
	return r.db.Save(p).Error
}

func (r *petRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Pet{}, id).Error
}
