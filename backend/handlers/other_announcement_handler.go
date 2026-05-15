package handlers

import (
	"bongthom_backend/data"
	"github.com/gofiber/fiber/v3"
)

func GetOtherAnnouncementList(c fiber.Ctx) error {
	return c.JSON(data.GetOtherAnnouncementList)
}