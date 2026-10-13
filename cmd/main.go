// @title E-Commerce Al-Hasra API
// @version 1.0
// @description API backend layanan e-commerce Sekolah Al-Hasra.
// @host localhost:8000
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import "github.com/jakewd7/internal/app"

func main() {
	app := app.NewApp()
	app.Run()
}
