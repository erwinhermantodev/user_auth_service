package repository

type User struct {
	ID          int
	PhoneNumber string
	FullName    string
	Password    string
}

type LoginInfo struct {
    PhoneNumber string `json:"phone_number"`
    Password    string `json:"password"`
}