package services

import "github.com/Venukishore-R/CODECRAFT_BW_01/internal/models"

type User struct {
	UserModel *models.UserModel
}

type UserService interface {
	CreateUser(user *models.User) *models.User
	GetAllUsersService() ([]*models.User, error)
	GetOneUserService(email string) (*models.User, error)
	UpdateUserService(user *models.User, email string) (*models.User, error)
	DeleteUserService(email string) error
	CreateUUID() string
}

func NewUserService(userModel *models.UserModel) *User {
	return &User{
		UserModel: userModel,
	}
}
