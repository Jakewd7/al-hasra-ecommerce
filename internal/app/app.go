package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jakewd7/internal/config"
	"github.com/jakewd7/internal/docs"
	"github.com/jakewd7/internal/middleware"
	"github.com/jakewd7/internal/routes"
	seeders "github.com/jakewd7/internal/seeder"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	config.LoadConfig()

	db := config.ConnectDatabase()

	// ================= SEEDER =================
	seeders.SeedDatabase(db)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(middleware.CORS())

	// ================= ROUTES =================
	api := r.Group("/api")
	routes.RegisterRoutes(api, db)

	// ================= SWAGGER =================
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &App{
		Router: r,
	}
}

func (a *App) Run() {
	port := config.Get("APP_PORT", "8080")
	log.Println("🚀 server running on", port)
	a.Router.Run("localhost:" + port)
}
