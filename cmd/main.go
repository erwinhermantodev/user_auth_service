package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/erwinhermantodev/user_auth_service/handler"
	"github.com/erwinhermantodev/user_auth_service/repository"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

const (
    host     = "localhost"
    port     = 5435
    user     = "myuser"
    password = "mypassword"
    dbname   = "auth_user_db"
)

func connectionString() string {
    return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
}

func main() {
	db, err := sql.Open("postgres", connectionString())
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
