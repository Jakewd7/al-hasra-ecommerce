package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jakewd7/internal/middleware"
	"github.com/jakewd7/internal/modules/users"
	"gorm.io/gorm"

	// [AUTO-IMPORT]
	"github.com/jakewd7/internal/modules/categories"
	"github.com/jakewd7/internal/modules/orders"
	"github.com/jakewd7/internal/modules/products"
	"github.com/jakewd7/internal/modules/sandboxes"
)

func RegisterRoutes(publicGroup *gin.RouterGroup, db *gorm.DB) {
	protectedGroup := publicGroup.Group("")
	protectedGroup.Use(middleware.AuthMiddleware())

	// hybrid
	users.RegisterRoutes(publicGroup, protectedGroup, db)

	//auth needed
	// roles.RegisterRoutes(protectedGroup, db)

	// [AUTO-REGISTER]
	sandboxes.RegisterRoutes(publicGroup, db)
	orders.RegisterRoutes(publicGroup, protectedGroup, db)
	products.RegisterRoutes(publicGroup, protectedGroup, db)
	categories.RegisterRoutes(publicGroup, protectedGroup, db)

}
