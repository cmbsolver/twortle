package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupUIRoutes initializes the views/UI related routes
func SetupUIRoutes(app *fiber.App) {
	// Home Page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Home",
		})
	})

	app.Get("/solver", func(c *fiber.Ctx) error {
		return c.Render("solver", fiber.Map{
			"Title": "Solver",
		})
	})

	app.Get("/drawtool", func(c *fiber.Ctx) error {
		return c.Render("drawtool", fiber.Map{
			"Title": "Drawing Tool",
		})
	})

	app.Get("/play", func(c *fiber.Ctx) error {
		return c.Render("play", fiber.Map{
			"Title": "Play Game",
		})
	})

}
