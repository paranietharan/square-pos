package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection
func Connect() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file, %s", err)
	}

	// Get database configuration details from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var errOpen error
	DB, errOpen = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errOpen != nil {
		log.Fatalf("Could not connect to the database: %v", errOpen)
	}

	fmt.Println("Successfully connected to PostgreSQL database successfully.........")
	return DB
}

func Disconnect() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Could not get raw DB object: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Error closing the database connection: %v", err)
	}

	fmt.Println("Database connection closed!")
}
