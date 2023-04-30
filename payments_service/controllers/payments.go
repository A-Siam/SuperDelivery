package controllers

import (
	"github.com/A-Siam/super_delivery/payments_service/data/models"
	"github.com/A-Siam/super_delivery/payments_service/services"
	"github.com/gofiber/fiber/v2"
)

func AddPayment(c *fiber.Ctx) error {
	var details models.PaymentDetails
	if err := c.BodyParser(&details); err != nil {
		return err
	}
	if err := services.DoPayment(details); err != nil {
		return err
	}
	return nil
}

func RevertPayment(c *fiber.Ctx) error {
	var event models.Event
	if err := c.BodyParser(&event); err != nil {
		return err
	}
	if err := services.RevertPayment(event); err != nil {
		return err
	}
	return nil
}
