package routes

import (
	"twortle/api"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Post("/search", api.SearchWordsHandler)
	apiGroup.Post("/draw", api.DrawPatternHandler)
}
