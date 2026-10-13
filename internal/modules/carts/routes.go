package carts

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	cartRoutes := r.Group("/carts")
	{
		cartRoutes.POST("", handler.AddToCart)
		cartRoutes.GET("", handler.GetCarts)
		cartRoutes.DELETE("/:id", handler.DeleteCart)
	}
}
