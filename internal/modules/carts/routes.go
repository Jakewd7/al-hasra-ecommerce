package carts

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	cart := r.Group("/carts")
	{
		cart.POST("", handler.CreateCart)
		cart.GET("", handler.GetAll)
		cart.GET("/:id", handler.GetByID)
		cart.PUT("/:id", handler.Update)
		cart.DELETE("/:id", handler.Delete)
	}
}
