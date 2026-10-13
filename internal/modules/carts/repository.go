package carts

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Repository interface {
	Create(data *Cart) error
	FindAll() ([]Cart, error)
	FindByID(id uuid.UUID) (*Cart, error)
	Update(data *Cart) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(data *Cart) error {
	return r.db.Create(data).Error
}

func (r *repository) FindAll() ([]Cart, error) {
	var results []Cart
	err := r.db.Find(&results).Error
	return results, err
}

func (r *repository) FindByID(id uuid.UUID) (*Cart, error) {
	var result Cart
	err := r.db.First(&result, "id = ?", id).Error
	return &result, err
}

func (r *repository) Update(data *Cart) error {
	return r.db.Save(data).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Cart{}, "id = ?", id).Error
}
