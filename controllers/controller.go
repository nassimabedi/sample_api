package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math"
	"sampleApi/database"
	"sampleApi/model"
	"sampleApi/pagination"
	"strconv"
	"strings"
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
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could user create user", "data": err})
	}

	// Return the created user
	return c.JSON(fiber.Map{"status": "success", "message": "Created User", "data": user})

}

// GetUser func one user by ID
// @Description Get one user by ID
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @router /api/user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	// Read the param userId
	id := c.Params("userId")

	// Find the user with the given Id
	db.Find(&user, "id = ?", id)

	// If no such user present return an error
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	// Return the user with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": user})
}

// UpdateUser update a user by ID
// @Description Update a User by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param first_name body string true "FirstName"
// @Param last_name body string true "LastName"
// @Param user_name body string true "UserName"
// @Success 200 {object} model.User
// @router /api/user/{id} [post]
func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		UserName  string `json:"user_name"`
	}
	db := database.DB
	var user model.User

	// Read the param userId
	id := c.Params("userId")

	// Find the user with the given Id
	db.Find(&user, "id = ?", id)

	// If no such user present return an error
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the user
	user.FirstName = updateUserData.FirstName
	user.LastName = updateUserData.LastName
	user.UserName = updateUserData.UserName

	// Save the Changes
	db.Save(&user)

	// Return the updated user
	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": user})
}

// DeleteUser delete a User by ID
// @Description Delete a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @router /api/User/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	// Read the param userId
	id := c.Params("userId")

	// Find the user with the given Id
	db.Find(&user, "id = ?", id)

	// If no such user present return an error
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	// Delete the user and return error if encountered
	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted User"})
}

func paginate(value interface{}, pagination *pagination.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

// GetUsers func gets all existing users
// @Description Get all existing users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @router /user [get]
func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	// var pagination *pagination
	var pagination pagination.Pagination
	var pageVar, limitVar int
	var err error

	page := strings.TrimSpace(c.Query("page"))
	limit := strings.TrimSpace(c.Query("limit"))
	sort := strings.TrimSpace(c.Query("sort"))

	if len(page) > 0 {
		pageVar, err = strconv.Atoi(page)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Error in page param", "data": nil})
		}
	} else {
		pageVar = 1
	}

	if len(limit) > 0 {
		limitVar, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Error in page param", "data": nil})
		}
	} else {
		limitVar = 10
	}

	if len(sort) > 0 {
		pagination.Sort = sort
	}

	pagination.Page = pageVar
	pagination.Limit = limitVar

	fmt.Println("=====page", pageVar, limitVar)
	fmt.Println(string(c.Request().URI().QueryString()))

	// find all users in the database
	db.Scopes(paginate(users, &pagination, db)).Find(&users)
	pagination.Rows = users

	// If no user is present return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No users present", "data": nil})
	}

	// Else return users
	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": pagination.Rows})
}
