package repository

type User struct {
	ID          int
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Password    string
}

type LoginInfo struct {
    PhoneNumber string `json:"phone_number"`
    Password    string `json:"password"`
}