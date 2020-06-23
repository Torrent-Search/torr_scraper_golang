package controller

import (
	"os"

	"github.com/gofiber/fiber"
	models "github.com/scraper_v2/models"
)

func AppUpdateController(fibCon *fiber.Ctx) {
	var update models.Update
	update.Version = os.Getenv("APP_VERSION")
	update.Link = os.Getenv("LINK")
	fibCon.Status(200).JSON(update)
}
