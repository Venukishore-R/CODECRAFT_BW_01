package services

import (
	"fmt"
	"regexp"

	"github.com/Venukishore-R/CODECRAFT_BW_01/internal/models"
	"github.com/google/uuid"
)

func (s *User) CreateUserService(user *models.User) (*models.User, error) {
	if !s.IsValidEmail(user.Email) {
		return nil, fmt.Errorf("invalid email: %v", user.Email)
	}
	return s.UserModel.CreateUser(user)
}

func (s *User) GetAllUsersService() ([]*models.User, error) {
	return s.UserModel.GetAllUsers()
}

func (s *User) GetOneUserService(email string) (*models.User, error) {
	return s.UserModel.GetOne(email)
}

func (s *User) UpdateUserService(user *models.User, email string) (*models.User, error) {
	return s.UserModel.UpdateUser(user, email)
}

func (s *User) DeleteUserService(email string) error {
	return s.UserModel.DeleteUser(email)
}

func (s *User) CreateUUID() string {
	id := uuid.New()
	return id.String()
}

func (s *User) IsValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	reg := regexp.MustCompile(emailPattern)

	return reg.MatchString(email)
}
