package products

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Repository interface {
	Create(data *Product) error
	FindAll() ([]Product, error)
	FindByID(id uuid.UUID) (*Product, error)
	Update(data *Product) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(data *Product) error {
	return r.db.Create(data).Error
}

func (r *repository) FindAll() ([]Product, error) {
	var results []Product
	err := r.db.Find(&results).Error
	return results, err
}

func (r *repository) FindByID(id uuid.UUID) (*Product, error) {
	var result Product
	err := r.db.First(&result, "id = ?", id).Error
	return &result, err
}

func (r *repository) Update(data *Product) error {
	return r.db.Save(data).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Product{}, "id = ?", id).Error
}
