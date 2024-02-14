package repository

type UserRepositoryInterface interface {
	CreateUser(user *User) (int, error)
	FindUserByPhone(phone string) (*User, error)
	IncrementLoginCount(userID int) error
	GetUserByID(userID int) (*User, error)
	UpdateUserProfile(user *User) error
}