package orders

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	publicOrders := public.Group("/orders")
	{
		publicOrders.POST("/checkout", handler.Checkout)
		publicOrders.GET("/:id", handler.GetOrderByID)
	}

	protectedOrders := protected.Group("/orders")
	{
		protectedOrders.GET("", handler.GetOrders)
	}
}
