package main

import (
	"twortle/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".tmpl")

	// Pass the engine to the Fiber app
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// Setup Routes
	app.Static("/assets", "./assets")
	routes.RegisterAPIRoutes(app)
	routes.SetupUIRoutes(app)

	app.Listen(":3000")
}
