package products

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Product, error)
	FindByID(id string) (*Product, error)
	FindBySlug(slug string) (*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *repository) FindByID(id string) (*Product, error) {
	var product Product
	err := r.db.First(&product, "id = ?", id).Error
	return &product, err
}

func (r *repository) FindBySlug(slug string) (*Product, error) {
	var product Product
	err := r.db.First(&product, "slug = ?", slug).Error
	return &product, err
}

func (r *repository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *repository) Update(product *Product) error {
	return r.db.Save(product).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&Product{}, "id = ?", id).Error
}
