package handlers

import (
	"bongthom_backend/data"
	"github.com/gofiber/fiber/v3"
)

func GetQuickLinksList(c fiber.Ctx) error {
	return c.JSON(data.GetQuickLinksList)
}