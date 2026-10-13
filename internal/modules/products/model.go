package products

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time 
	UpdatedAt   time.Time 
}

func (m *Product) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
