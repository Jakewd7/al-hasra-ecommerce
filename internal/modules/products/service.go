package products

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(req CreateProductRequest) (*Product, error) {
	data := &Product{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
	}
	err := s.repo.Create(data)
	return data, err
}

func (s *Service) GetAll() ([]Product, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id uuid.UUID, req UpdateProductRequest) (*Product, error) {
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
