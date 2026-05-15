package handlers

import (
	"bongthom_backend/data"
	"github.com/gofiber/fiber/v3"
)

func GetLatestJobList(c fiber .Ctx) error {
	return c.JSON(data.LatestJobList)
}