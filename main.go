package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"sampleApi/database"
	"sampleApi/routes"
)

func main() {
	// Start a new fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Listen on PORT 3000
	log.Fatal(app.Listen(":3000"))

}
