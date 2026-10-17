package sandboxes

import (
	"github.com/jakewd7/internal/modules/orders"
	"gorm.io/gorm"
)

type Repository interface {
	FindByVANumber(vaNumber string) (*orders.Payment, error)
	GetOrderWithDetails(orderID string) (*orders.Order, error)
	GetDB() *gorm.DB
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

func (r *repository) FindByVANumber(vaNumber string) (*orders.Payment, error) {
	var payment orders.Payment
	err := r.db.Where("va_number = ?", vaNumber).First(&payment).Error
	return &payment, err
}

func (r *repository) GetOrderWithDetails(orderID string) (*orders.Order, error) {
	var order orders.Order
	err := r.db.Preload("OrderItems.Product").Preload("Payment").First(&order, "id = ?", orderID).Error
	if err != nil {
		err = r.db.Preload("OrderItems.Product").First(&order, "id = ?", orderID).Error
	}
	return &order, err
}
