package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	userPublic := public.Group("/users")
	{
		userPublic.POST("/login", handler.Login)
		userPublic.POST("/refresh", handler.Refresh)
		userPublic.POST("/register", handler.Register)
	}

	userPrivate := protected.Group("/users")
	{
		// --- PERBAIKAN DI SINI ---
		// Ganti "/" menjadi "" agar route-nya menjadi "/api/users" (tanpa slash di akhir)
		userPrivate.GET("", handler.GetUsers)

		// Route parameter (:id) sudah benar pakai slash karena diawali parameter
		userPrivate.PUT("/:id", handler.UpdateUser)
		userPrivate.DELETE("/:id", handler.DeleteUser)

		// Route statis lain
		// userPrivate.POST("/register", handler.Register)
		userPrivate.POST("/logout", handler.Logout)
	}
}
