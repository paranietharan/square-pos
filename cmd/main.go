package main

import (
	"log"
	"square-pos/pkg/config"
	"square-pos/pkg/router"
)

func main() {
	log.Printf("Using JWT Secret: %s", config.Envs.JWTSecret)

	// Connect to the database
	db := config.Connect()
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	// Start the server
	router.StartServer(db)
	log.Println("Server is running on port 8080...")

	// Keep the server running
	// The select key word is used to run the go rountines indefinately
	// It helps to keep the server up and running other wise it crash the server
	// refer the screen shot in the Screenshot_01
	select {}
}
