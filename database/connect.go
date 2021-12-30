package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sampleApi/model"
	"strconv"
)

// Declare the variable for the database
var DB *gorm.DB

// convert string to int from env
func getenvInt(str string) (int, error) {
	v, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return v, nil
}

// ConnectDB connect to db
func ConnectDB() {
	var err error

	port, err := getenvInt(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"),
		port, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")

	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}
