package orders

import (
	"github.com/jakewd7/internal/modules/products"
	"gorm.io/gorm"
)

type Repository interface {
	CreateOrder(tx *gorm.DB, order *Order) error
	FindAll() ([]Order, error)
	FindByID(id string) (*Order, error)
	GetDB() *gorm.DB
	FindProductByID(id string) (*products.Product, error)
	UpdateProductStock(tx *gorm.DB, product *products.Product) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetDB() *gorm.DB {
	return r.db
}

func (r *repository) CreateOrder(tx *gorm.DB, order *Order) error {
	return tx.Create(order).Error
}

func (r *repository) FindAll() ([]Order, error) {
	var orders []Order
	err := r.db.Preload("OrderItems.Product").Find(&orders).Error
	return orders, err
}

func (r *repository) FindByID(id string) (*Order, error) {
	var order Order
	err := r.db.Preload("OrderItems.Product").First(&order, "id = ?", id).Error
	return &order, err
}

func (r *repository) FindProductByID(id string) (*products.Product, error) {
	var product products.Product
	err := r.db.First(&product, "id = ?", id).Error
	return &product, err
}

func (r *repository) UpdateProductStock(tx *gorm.DB, product *products.Product) error {
	return tx.Save(product).Error
}
