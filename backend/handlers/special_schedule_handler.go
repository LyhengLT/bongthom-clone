package handlers

import (
	"bongthom_backend/data"
	"github.com/gofiber/fiber/v3"
)

func GetSpecialScheduleList(c fiber.Ctx) error {
	return c.JSON(data.GetSpecialSchedulesList)
}