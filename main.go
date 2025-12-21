package main

import (
	"fmt"
	"twortle/pkg/db"
	"twortle/pkg/db/tables"
	"twortle/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create and load the database
	dbConn, _ := db.InitSQLiteConnection()

	migrateErr := db.InitSQLiteTables(dbConn)
	if migrateErr != nil {
		fmt.Printf("Error initializing database tables: %v\n", migrateErr)
	}

	if tables.GetWordCount(dbConn) == 0 {
		fmt.Println("No words found in database, loading words.txt...")
		tables.LoadFile(dbConn, "./assets/words.txt")
	}

	closeError := db.CloseConnection(dbConn)
	if closeError != nil {
		fmt.Printf("Error closing database connection: %v\n", closeError)
	}

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
