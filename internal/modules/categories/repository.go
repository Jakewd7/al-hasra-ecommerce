package categories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Category, error)
	FindByID(id uuid.UUID) (Category, error)
	Create(category *Category) error
	Update(category *Category) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *repository) FindByID(id uuid.UUID) (Category, error) {
	var category Category
	err := r.db.First(&category, "id = ?", id).Error
	return category, err
}

func (r *repository) Create(category *Category) error {
	return r.db.Create(category).Error
}

func (r *repository) Update(category *Category) error {
	return r.db.Save(category).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Category{}, "id = ?", id).Error
}
