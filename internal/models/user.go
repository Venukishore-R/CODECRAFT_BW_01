package models

type User struct {
	Id    string
	Name  string
	Email string
	Age   int
}

type UserModel struct {
	Users []*User
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetAllUsers() ([]*User, error)
	GetOne(email string) (*User, error)
	UpdateUser(user *User, email string) (*User, error)
	DeleteUser(email string) error
	CheckUserExists(email string) bool
}

func NewUserModel(users []*User) *UserModel {
	return &UserModel{
		Users: users,
	}
}
