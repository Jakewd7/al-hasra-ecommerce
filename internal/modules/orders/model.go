package orders

import (
	"time"

	"github.com/google/uuid"
	"github.com/jakewd7/internal/modules/products"
	"gorm.io/gorm"
)

type Order struct {
	ID          string      `gorm:"type:varchar(100);primaryKey"`
	OrderNumber string      `gorm:"type:varchar(50);unique;not null"`
	TotalAmount float64     `gorm:"type:decimal(12,2);not null;default:0"`
	Status      string      `gorm:"type:varchar(30);default:'PENDING'"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID"`
	Payment     *Payment    `gorm:"foreignKey:OrderID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
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

type Payment struct {
	ID            string     `gorm:"type:varchar(100);primaryKey"`
	OrderID       string     `gorm:"type:varchar(100);not null;unique"`
	PaymentMethod string     `gorm:"type:varchar(50);not null;default:'BCA_VA'"`
	VANumber      string     `gorm:"type:varchar(50);not null;unique"`
	Amount        float64    `gorm:"type:decimal(12,2);not null;default:0"`
	Status        string     `gorm:"type:varchar(30);default:'PENDING'"`
	ExpiredAt     time.Time  `gorm:"not null"`
	PaidAt        *time.Time `gorm:"type:datetime"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
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

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

type CheckoutItemRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
}

type CheckoutRequest struct {
	Items []CheckoutItemRequest `json:"items" binding:"required,gt=0"`
}

type OrderItemResponse struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}

type PaymentResponse struct {
	ID            string     `json:"payment_id"`
	PaymentMethod string     `json:"payment_method"`
	VANumber      string     `json:"va_number"`
	Amount        float64    `json:"amount"`
	Status        string     `json:"status"`
	ExpiredAt     time.Time  `json:"expired_at"`
	ExpiresIn     int64      `json:"expires_in"`
	PaidAt        *time.Time `json:"paid_at,omitempty"`
}

type OrderResponse struct {
	ID          string              `json:"order_id"`
	OrderNumber string              `json:"order_number"`
	TotalAmount float64             `json:"total_amount"`
	Status      string              `json:"status"`
	OrderItems  []OrderItemResponse `json:"items"`
	Payment     *PaymentResponse    `json:"payment,omitempty"`
	CreatedAt   time.Time           `json:"created_at"`
}
