package sandboxes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(public *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	sandboxRoutes := public.Group("/sandbox")
	{
		sandboxRoutes.GET("/va/:va_number", handler.GetPaymentByVA)
		sandboxRoutes.POST("/pay", handler.PaySandbox)
	}
}
