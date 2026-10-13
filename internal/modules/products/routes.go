package products

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	publicProducts := public.Group("/products")
	{
		publicProducts.GET("", handler.GetProducts)
		publicProducts.GET("/:id", handler.GetProductByID)
	}

	protectedProducts := protected.Group("/products")
	{
		protectedProducts.POST("", handler.CreateProduct)
		protectedProducts.PUT("/:id", handler.UpdateProduct)
		protectedProducts.DELETE("/:id", handler.DeleteProduct)
	}
}
