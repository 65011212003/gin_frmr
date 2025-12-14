package main

import (
	"gin_frmr/database"
	"gin_frmr/routes"
	"log"
)

func main() {
	// Initialize database
	database.Connect()

	// Setup router
	r := routes.SetupRouter()

	// Start server
	log.Println("Server starting on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
