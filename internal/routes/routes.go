package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jakewd7/internal/middleware"
	"github.com/jakewd7/internal/modules/roles"
	"github.com/jakewd7/internal/modules/users"
	"gorm.io/gorm"

	// [AUTO-IMPORT]
	"github.com/jakewd7/internal/modules/orders"
	"github.com/jakewd7/internal/modules/carts"
	"github.com/jakewd7/internal/modules/categories"
	"github.com/jakewd7/internal/modules/products"
)

func RegisterRoutes(publicGroup *gin.RouterGroup, db *gorm.DB) {
	protectedGroup := publicGroup.Group("")
	protectedGroup.Use(middleware.AuthMiddleware())

	// hybrid
	users.RegisterRoutes(publicGroup, protectedGroup, db)

	//auth needed
	roles.RegisterRoutes(protectedGroup, db)

	// [AUTO-REGISTER]
	orders.RegisterRoutes(protectedGroup, db)
	carts.RegisterRoutes(protectedGroup, db)
	products.RegisterRoutes(protectedGroup, publicGroup, db)
	categories.RegisterRoutes(protectedGroup, publicGroup, db)

}
