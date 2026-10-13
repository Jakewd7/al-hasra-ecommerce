package products

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	product := r.Group("/products")
	{
		product.POST("", handler.CreateProduct)
		product.GET("", handler.GetAll)
		product.GET("/:id", handler.GetByID)
		product.PUT("/:id", handler.Update)
		product.DELETE("/:id", handler.Delete)
	}
}
