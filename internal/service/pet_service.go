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

func (s *PetService) List() ([]domain.Pet, error) {
	return s.repo.FindAll()
}

func (s *PetService) Get(id uint) (*domain.Pet, error) {
	return s.repo.FindByID(id)
}

func (s *PetService) Update(id uint, in *domain.Pet) (*domain.Pet, error) {
	p, err := s.repo.FindByID(id)
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

func (s *PetService) Delete(id uint) error {
	return s.repo.Delete(id)
}
