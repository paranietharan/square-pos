package main

import (
	"fmt"
	"log"
	"net/http"
	"square-pos/pkg/config"
	"square-pos/pkg/router"
)

func main() {
	db := config.Connect()
	defer config.Disconnect()

	// err := db.AutoMigrate(&models.User{}, &models.Ride{}, &models.UserRides{})
	// if err != nil {
	// 	log.Fatalf("Error migrating schema: %v", err)
	// }
	// log.Println("Database tables migrated successfully!")

	r := router.InitializeRoutes(db)

	port := ":8080"
	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
