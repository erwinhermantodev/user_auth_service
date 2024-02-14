package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/erwinhermantodev/user_auth_service/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo}
}

func validateUserInput(user repository.User) error {
	if len(user.PhoneNumber) < 10 || len(user.PhoneNumber) > 13 && !strings.HasPrefix(user.PhoneNumber, "+62") {
		return errors.New("phone number must be 10-13 characters long and start with '+62'")
	}

	if len(user.FullName) < 3 || len(user.FullName) > 60 {
		return errors.New("full name must be 3-60 characters long")
	}

	if len(user.Password) < 6 || len(user.Password) > 64 {
		return errors.New("password must be 6-64 characters long")
	}

	hasUppercase := false
	hasNumber := false
	hasSpecial := false

	for _, char := range user.Password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUppercase = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' || char == '^' || char == '&' || char == '*':
			hasSpecial = true
		}
	}

	if !hasUppercase || !hasNumber || !hasSpecial {
		return errors.New("password must contain at least 1 uppercase letter, 1 number, and 1 special character")
	}

	return nil
}

func (h *UserHandler) Register(c echo.Context) error {
	var user repository.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	
	if err := validateUserInput(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	
	userID, err := h.repo.CreateUser(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]int{"id": userID})
}

func validateLoginInfo(loginInfo repository.LoginInfo) error {
	if len(loginInfo.PhoneNumber) < 10 || len(loginInfo.PhoneNumber) > 13 || !strings.HasPrefix(loginInfo.PhoneNumber, "+62") {
		return errors.New("phone number must be 10-13 characters long and start with '+62'")
	}

	if len(loginInfo.Password) < 6 || len(loginInfo.Password) > 64 {
		return errors.New("password must be 6-64 characters long")
	}

	return nil
}

func (h *UserHandler) Login(c echo.Context) error {
	var loginInfo struct {
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
	if err := c.Bind(&loginInfo); err != nil {
		return err
	}

	if err := validateLoginInfo(loginInfo); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

	user, err := h.repo.FindUserByPhone(loginInfo.PhoneNumber)
	fmt.Println("===============")
	fmt.Println("err")
	fmt.Println(err)
	fmt.Println("===============")
	if err != nil {
		return err
	}
	if user == nil {
		return echo.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return err
	}

	if err := h.repo.IncrementLoginCount(user.ID); err != nil {
		c.Logger().Errorf("Error incrementing login count: %v", err)
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func (h *UserHandler) GetProfile(c echo.Context) error {
	userID := c.Get("user").(int)

	user, err := h.repo.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":       user.ID,
		"phone":    user.PhoneNumber,
		"full_name": user.FullName,
	})
}

func (h *UserHandler) UpdateProfile(c echo.Context) error {
	userID := c.Get("user").(int)

	var updateUser repository.User
	if err := c.Bind(&updateUser); err != nil {
		return err
	}

	if err := validateUpdateUserFields(updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if updateUser.ID != userID {
		return echo.ErrForbidden
	}

	if err := h.repo.UpdateUserProfile(&updateUser); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Profile updated successfully")
}

func validateUpdateUserFields(updateUser repository.User) error {
	if updateUser.FullName == "" {
		return errors.New("full name cannot be empty")
	}

	if len(updateUser.PhoneNumber) < 10 || len(updateUser.PhoneNumber) > 13 || !strings.HasPrefix(updateUser.PhoneNumber, "+62") {
		return errors.New("phone number must be 10-13 characters long and start with '+62'")
	}

	return nil
}
