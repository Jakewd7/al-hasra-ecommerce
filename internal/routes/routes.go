package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jakewd7/internal/middleware"
	"github.com/jakewd7/internal/modules/roles"
	"github.com/jakewd7/internal/modules/users"
	"gorm.io/gorm"
	// [AUTO-IMPORT]
)

func RegisterRoutes(publicGroup *gin.RouterGroup, db *gorm.DB) {
	protectedGroup := publicGroup.Group("")
	protectedGroup.Use(middleware.AuthMiddleware())

	// hybrid
	users.RegisterRoutes(publicGroup, protectedGroup, db)

	//auth needed
	roles.RegisterRoutes(protectedGroup, db)

	// [AUTO-REGISTER]

}
