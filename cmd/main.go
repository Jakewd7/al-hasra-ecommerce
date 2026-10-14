// @title E-Commerce Al-Hasra API
// @version 1.0
// @description API backend layanan e-commerce Sekolah Al-Hasra.
// @host serverstaging-vps-5e6db201b3.asiahub.id
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
