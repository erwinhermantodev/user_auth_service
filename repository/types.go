package repository

import "time"

type User struct {
	ID          int       `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	FullName    string    `json:"full_name"`
	Password    string 	  `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type LoginInfo struct {
    PhoneNumber string `json:"phone_number"`
    Password    string `json:"password"`
}