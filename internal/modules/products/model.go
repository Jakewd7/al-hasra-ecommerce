package products

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string  `gorm:"type:varchar(100);primaryKey"`
	CategoryID  string  `gorm:"type:varchar(100);not null"`
	Name        string  `gorm:"type:varchar(150);not null"`
	Slug        string  `gorm:"type:varchar(150);uniqueIndex;not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(12,2);not null;default:0"`
	Stock       int     `gorm:"not null;default:0"`
	ImageURL    string  `gorm:"type:varchar(255)"`
	IsActive    bool    `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

type CreateProductRequest struct {
	CategoryID  string  `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required" example:"Sepatu Olahraga"`
	Description string  `json:"description" example:"Sepatu olahraga nyaman dipakai"`
	Price       float64 `json:"price" binding:"required,gt=0" example:"150000"`
	Stock       int     `json:"stock" binding:"required,gte=0" example:"10"`
	ImageURL    string  `json:"image_url" example:"https://example.com/image.jpg"`
}

type UpdateProductRequest struct {
	CategoryID  string   `json:"category_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *int     `json:"stock"`
	ImageURL    string   `json:"image_url"`
	IsActive    *bool    `json:"is_active"`
}

type ProductResponse struct {
	ID          string    `json:"id"`
	CategoryID  string    `json:"category_id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
