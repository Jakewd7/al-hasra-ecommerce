package categories

import "github.com/google/uuid"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() ([]Category, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (Category, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Create(req CreateCategoryRequest) (Category, error) {
	category := Category{
		Name: req.Name,
		Slug: req.Slug,
	}
	err := s.repo.Create(&category)
	return category, err
}

func (s *Service) Update(id uuid.UUID, req UpdateCategoryRequest) (Category, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return category, err
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Slug != "" {
		category.Slug = req.Slug
	}

	err = s.repo.Update(&category)
	return category, err
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
