package handlers

import (
	"bongthom_backend/data"
	"github.com/gofiber/fiber/v3"
)

func GetRSSList(c fiber.Ctx) error {
	return c.JSON(data.GetRSSList)
}