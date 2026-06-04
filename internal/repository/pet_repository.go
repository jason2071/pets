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

func (r *petRepository) FindAllByOwner(ownerID uint) ([]domain.Pet, error) {
	var pets []domain.Pet
	err := r.db.Where("owner_id = ?", ownerID).Order("id").Find(&pets).Error
	return pets, err
}

func (r *petRepository) FindByIDAndOwner(id, ownerID uint) (*domain.Pet, error) {
	var p domain.Pet
	if err := r.db.Where("id = ? AND owner_id = ?", id, ownerID).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *petRepository) Update(p *domain.Pet) error {
	return r.db.Save(p).Error
}

func (r *petRepository) DeleteByOwner(id, ownerID uint) error {
	res := r.db.Where("owner_id = ?", ownerID).Delete(&domain.Pet{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
