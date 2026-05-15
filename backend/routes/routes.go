package routes

import (
	"bongthom_backend/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/latest-job", handlers.GetLatestJobList)
	app.Get("/latest-classified", handlers.GetLatestClassifiedList)
	app.Get("/organization-type", handlers.GetOrganizationTypeList)
	app.Get("/other-announcement", handlers.GetOtherAnnouncementList)
	app.Get("/quick-link", handlers.GetQuickLinksList)
	app.Get("/rss", handlers.GetRSSList)
	app.Get("/special-schedule", handlers.GetSpecialScheduleList)
}