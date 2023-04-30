package main

import (
	"github.com/A-Siam/super_delivery/payments_service/ports"
	"github.com/A-Siam/super_delivery/payments_service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	ports.InitializeEnv()

	app := fiber.New()
	routes.SetupRoutes(app)
	serverPort := viper.GetString("server.port")
	app.Listen(":" + serverPort)
}
