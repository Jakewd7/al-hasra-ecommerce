package roles

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	roleRepo := NewRepository(db)
	roleService := NewService(roleRepo)
	roleHandler := NewHandler(roleService)

	roles := r.Group("/roles")
	{
		roles.GET("", roleHandler.GetRoles)
		roles.POST("", roleHandler.CreateRole)
		roles.PUT("/:id", roleHandler.UpdateRole)
		roles.DELETE("/:id", roleHandler.DeleteRole)
	}

	permission := r.Group("/permissions")
	{
		permission.GET("", roleHandler.GetPermissions)
	}

}
