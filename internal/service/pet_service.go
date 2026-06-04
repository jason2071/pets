package service

import (
	"errors"

	"github.com/jason2071/pets/internal/domain"
)

// ErrPetNotFound is returned when a pet does not exist.
var ErrPetNotFound = errors.New("pet not found")

// PetService holds business logic for pets.
type PetService struct {
	repo domain.PetRepository
}

// NewPetService constructs a PetService.
func NewPetService(repo domain.PetRepository) *PetService {
	return &PetService{repo: repo}
}

func (s *PetService) Create(p *domain.Pet) error {
	return s.repo.Create(p)
}

func (s *PetService) List(ownerID uint) ([]domain.Pet, error) {
	return s.repo.FindAllByOwner(ownerID)
}

func (s *PetService) Get(id, ownerID uint) (*domain.Pet, error) {
	return s.repo.FindByIDAndOwner(id, ownerID)
}

// Update modifies a pet the caller owns. Ownership is fixed; only mutable
// fields are changed.
func (s *PetService) Update(id, ownerID uint, in *domain.Pet) (*domain.Pet, error) {
	p, err := s.repo.FindByIDAndOwner(id, ownerID)
	if err != nil {
		return nil, err
	}
	p.Name = in.Name
	p.Species = in.Species
	p.Breed = in.Breed
	p.Age = in.Age
	if err := s.repo.Update(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PetService) Delete(id, ownerID uint) error {
	return s.repo.DeleteByOwner(id, ownerID)
}
