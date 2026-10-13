package carts

import (
	"time"

	"github.com/google/uuid"
	"github.com/jakewd7/internal/modules/products"
	"gorm.io/gorm"
)

type Cart struct {
	ID        string           `gorm:"type:varchar(100);primaryKey"`
	UserID    *string          `gorm:"type:varchar(100)"`
	ProductID string           `gorm:"type:varchar(100);not null"`
	Product   products.Product `gorm:"foreignKey:ProductID"`
	Quantity  int              `gorm:"not null;default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

type AddToCartRequest struct {
	UserID    *string `json:"user_id"`
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0" example:"1"`
}

type CartResponse struct {
	ID        string                   `json:"id"`
	UserID    *string                  `json:"user_id,omitempty"`
	ProductID string                   `json:"product_id"`
	Product   products.ProductResponse `json:"product"`
	Quantity  int                      `json:"quantity"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}
