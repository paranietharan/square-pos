package config

import (
	"fmt"
	"log"
	"os"
	"square-pos/pkg/types"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection
// TODO: Seperate DB functionality into diff package
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

	// Migration
	log.Println("Migration started.............")
	err = DB.AutoMigrate(&types.User{}, &types.Product{}, &types.Order{}, &types.OrderItem{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully!")
	log.Println("Successfully connected to PostgreSQL database successfully.........")
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

//-------------

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             GetEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   GetEnv("PORT", "8080"),
		DBUser:                 GetEnv("DB_USER", "root"),
		DBPassword:             GetEnv("DB_PASSWORD", "mypassword"),
		DBAddress:              fmt.Sprintf("%s:%s", GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "3306")),
		DBName:                 GetEnv("DB_NAME", "ecom"),
		JWTSecret:              GetEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

// Gets the env by key or fallbacks
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
