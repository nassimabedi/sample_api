package routes

import (
	"github.com/gofiber/fiber/v2"
	"sampleApi/controllers"
)

func SetupRoutes(router fiber.Router) {
	user := router.Group("/user")
	// Create a User
	user.Post("/", controllers.CreateUser)
	// Read all Users
	user.Get("/", controllers.GetUsers)
	// Read one User
	user.Get("/:userId", controllers.GetUser)
	// Update one User
	user.Put("/:userId", controllers.UpdateUser)
	// // // Delete one User
	user.Delete("/:userId", controllers.DeleteUser)
}
