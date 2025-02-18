package main

import (
	"log"
	"square-pos/pkg/config"
	"square-pos/pkg/router"
)

func main() {
	// Connect to the database
	db := config.Connect()
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	// Start the server
	router.StartServer(db)
	log.Println("Server is running on port 8080...")

	// Keep the server running
	select {}
}
