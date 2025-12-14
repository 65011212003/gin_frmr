package main

import (
	"log"

	"gin_frmr/internal/delivery/http/handler"
	"gin_frmr/internal/delivery/http/router"
	"gin_frmr/internal/infrastructure/database"
	"gin_frmr/internal/repository"
	"gin_frmr/internal/usecase"
)

func main() {
	// Initialize database
	db, err := database.NewSQLiteDB("app.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize layers (Dependency Injection)
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	// Setup router
	r := router.SetupRouter(userHandler)

	// Start server
	log.Println("Server starting on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
