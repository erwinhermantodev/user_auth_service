package main

import (
	"database/sql"
	"log"

	"github.com/erwinhermantodev/user_auth_service/handler"
	"github.com/erwinhermantodev/user_auth_service/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://user:password@localhost:5432/dbname?sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()
    e := echo.New()
    
    // Initialize repository
    repo := repository.NewUserRepository(db)

    // Initialize handlers
    userHandler := handler.NewUserHandler(repo)

    // Register routes
    e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)
	e.GET("/profile", userHandler.GetProfile)
	e.PUT("/profile", userHandler.UpdateProfile)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
