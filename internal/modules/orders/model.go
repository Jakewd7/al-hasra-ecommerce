package orders

import (
	"time"

	"github.com/google/uuid"
	"github.com/jakewd7/internal/modules/products"
	"gorm.io/gorm"
)

type Order struct {
	ID              string      `gorm:"type:varchar(100);primaryKey"`
	OrderNumber     string      `gorm:"type:varchar(50);unique;not null"`
	UserID          *string     `gorm:"type:varchar(100)"`
	CustomerName    string      `gorm:"type:varchar(150);not null"`
	CustomerPhone   string      `gorm:"type:varchar(20);not null"`
	TotalAmount     float64     `gorm:"type:decimal(12,2);not null;default:0"`
	Status          string      `gorm:"type:varchar(30);default:'PENDING'"`
	ShippingAddress string      `gorm:"type:text;not null"`
	OrderItems      []OrderItem `gorm:"foreignKey:OrderID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type OrderItem struct {
	ID        string           `gorm:"type:varchar(100);primaryKey"`
	OrderID   string           `gorm:"type:varchar(100);not null"`
	ProductID string           `gorm:"type:varchar(100);not null"`
	Product   products.Product `gorm:"foreignKey:ProductID"`
	Price     float64          `gorm:"type:decimal(12,2);not null"`
	Quantity  int              `gorm:"not null"`
	Subtotal  float64          `gorm:"type:decimal(12,2);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}
	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = uuid.New().String()
	}
	return nil
}

type CheckoutItemRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
}

type CheckoutRequest struct {
	UserID          *string               `json:"user_id"`
	CustomerName    string                `json:"customer_name" binding:"required" example:"John Doe"`
	CustomerPhone   string                `json:"customer_phone" binding:"required" example:"08123456789"`
	ShippingAddress string                `json:"shipping_address" binding:"required" example:"Jl. Raya Ciputat No. 12"`
	Items           []CheckoutItemRequest `json:"items" binding:"required,gt=0"`
}

type OrderItemResponse struct {
	ID        string                   `json:"id"`
	ProductID string                   `json:"product_id"`
	Product   products.ProductResponse `json:"product"`
	Price     float64                  `json:"price"`
	Quantity  int                      `json:"quantity"`
	Subtotal  float64                  `json:"subtotal"`
}

type OrderResponse struct {
	ID              string              `json:"id"`
	OrderNumber     string              `json:"order_number"`
	UserID          *string             `json:"user_id,omitempty"`
	CustomerName    string              `json:"customer_name"`
	CustomerPhone   string              `json:"customer_phone"`
	TotalAmount     float64             `json:"total_amount"`
	Status          string              `json:"status"`
	ShippingAddress string              `json:"shipping_address"`
	OrderItems      []OrderItemResponse `json:"order_items"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
}
