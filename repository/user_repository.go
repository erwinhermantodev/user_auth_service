package repository

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user *User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	var userID int
	err = ur.db.QueryRow("INSERT INTO users (phone_number, full_name, password) VALUES ($1, $2, $3) RETURNING id",
		user.PhoneNumber, user.FullName, hashedPassword).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (ur *userRepository) FindUserByPhone(phone string) (*User, error) {
	var user User
	err := ur.db.QueryRow("SELECT id, phone_number, full_name, password FROM users WHERE phone_number = $1", phone).
		Scan(&user.ID, &user.PhoneNumber, &user.FullName, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) IncrementLoginCount(userID int) error {
	_, err := ur.db.Exec("UPDATE users SET login_count = login_count + 1 WHERE id = $1", userID)
	return err
}

func (ur *userRepository) GetUserByID(userID int) (*User, error) {
	var user User
	err := ur.db.QueryRow("SELECT id, phone_number, full_name, password FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.PhoneNumber, &user.FullName, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) UpdateUserProfile(user *User) error {
	_, err := ur.db.Exec("UPDATE users SET full_name = $1 WHERE id = $2", user.FullName, user.ID)
	return err
}