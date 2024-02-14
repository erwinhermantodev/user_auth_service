package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erwinhermantodev/user_auth_service/generated"
	"github.com/erwinhermantodev/user_auth_service/handler"
	"github.com/erwinhermantodev/user_auth_service/middleware"
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
	e.GET("/profile", userHandler.GetProfile, middleware.JWTAuthMiddleware)
	e.PUT("/profile", userHandler.UpdateProfile, middleware.JWTAuthMiddleware)

	// Serve Swagger UI
    e.GET("/swagger", echoSwaggerMiddleware())

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

// echoSwaggerMiddleware serves the Swagger UI files
func echoSwaggerMiddleware() echo.HandlerFunc {
	return func(c echo.Context) error {
		spec, err := generated.GetSwagger()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, spec)
	}
}