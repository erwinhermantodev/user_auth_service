package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
        }

        tokenParts := strings.Split(tokenString, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header format")
        }

        token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
            // Here you should validate the token signature using your secret key or public key
            // Replace "your-secret-key" with your actual secret key
            return []byte("your-secret-key"), nil
        })
        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT token")
        }

        if !token.Valid {
            return echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT token")
        }

        // Extract user ID from the token and store it in the context
        claims := token.Claims.(jwt.MapClaims)
		userIDFloat64, ok := claims["user_id"].(float64)
        if !ok {
            return echo.NewHTTPError(http.StatusUnauthorized, "invalid user ID in JWT token")
        }
        userID := int(userIDFloat64)
        c.Set("userID", userID)

        return next(c)
    }
}
