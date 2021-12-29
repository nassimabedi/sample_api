package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // Adds some metadata fields to the table
	ID         uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	UserName   string    `json:"user_name"`
}
