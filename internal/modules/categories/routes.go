package categories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	category := r.Group("/categories")
	{
		category.POST("", handler.CreateCategory)
		category.GET("", handler.GetAll)
		category.GET("/:id", handler.GetByID)
		category.PUT("/:id", handler.Update)
		category.DELETE("/:id", handler.Delete)
	}
}
