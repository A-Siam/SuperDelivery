package routes

import (
	"github.com/A-Siam/super_delivery/payments_service/controllers"
	"github.com/gofiber/fiber/v2"
)

func setupPaymentRoutes(app *fiber.App) {
	app.Get("/payments", controllers.GetAllPayments)
	app.Post("/payments", controllers.AddPayment)
	app.Post("/payments/revert", controllers.RevertPayment)
}
