package carts

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(cart *Cart) error
	FindAll() ([]Cart, error)
	FindByID(id string) (*Cart, error)
	Update(cart *Cart) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(cart *Cart) error {
	return r.db.Create(cart).Error
}

func (r *repository) FindAll() ([]Cart, error) {
	var carts []Cart
	err := r.db.Preload("Product").Find(&carts).Error
	return carts, err
}

func (r *repository) FindByID(id string) (*Cart, error) {
	var cart Cart
	err := r.db.Preload("Product").First(&cart, "id = ?", id).Error
	return &cart, err
}

func (r *repository) Update(cart *Cart) error {
	return r.db.Save(cart).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&Cart{}, "id = ?", id).Error
}
