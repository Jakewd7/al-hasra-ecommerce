package carts

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(req CreateCartRequest) (*Cart, error) {
	data := &Cart{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
	}
	err := s.repo.Create(data)
	return data, err
}

func (s *Service) GetAll() ([]Cart, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*Cart, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id uuid.UUID, req UpdateCartRequest) (*Cart, error) {
	data, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	
	if req.Name != "" {
		data.Name = req.Name
	}
	data.Description = req.Description

	err = s.repo.Update(data)
	return data, err
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
