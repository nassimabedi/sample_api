package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"sampleApi/database"
	"sampleApi/model"
)


func CreateUser(c *fiber.Ctx) error {
	
	db := database.DB
	user := new(model.User)
	
	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the user
	user.ID = uuid.New()
	
	// Create the User and return error if encountered
	err = db.Create(&user).Error
	
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	
	// Return the created user
	return c.JSON(fiber.Map{"status": "success", "message": "Created User", "data": user})

	
}
