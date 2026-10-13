package categories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	publicCategories := public.Group("/products")
	{
		publicCategories.GET("", handler.GetAll)
		publicCategories.GET("/:id", handler.GetByID)
	}

	protectedCategories := protected.Group("/products")
	{
		protectedCategories.POST("", handler.CreateCategory)
		protectedCategories.PUT("/:id", handler.Update)
		protectedCategories.DELETE("/:id", handler.Delete)
	}
}
