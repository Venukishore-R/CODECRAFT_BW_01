package models

import "fmt"

func (u *UserModel) CreateUser(user *User) (*User, error) {
	if !u.CheckUserExists(user.Email) {
		u.Users = append(u.Users, user)
		return user, nil
	}

	return nil, fmt.Errorf("user already exists for email %s", user.Email)
}

func (u *UserModel) GetAllUsers() ([]*User, error) {
	if u.Users == nil {
		return nil, fmt.Errorf("no users found")
	}

	return u.Users, nil
}

func (u *UserModel) GetOne(email string) (*User, error) {
	for _, user := range u.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found for email %s", email)
}

func (u *UserModel) UpdateUser(user *User, email string) (*User, error) {
	for _, u := range u.Users {
		if u.Email == email {
			*u = *user
			return u, nil
		}
	}
	return nil, fmt.Errorf("user not found for email %s", email)
}

func (u *UserModel) DeleteUser(email string) error {
	for i := range u.Users {
		if u.Users[i].Email == email {
			u.Users = append(u.Users[:i], u.Users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user not found for email %s", email)
}

func (u *UserModel) CheckUserExists(email string) bool {
	for _, user := range u.Users {
		if user.Email == email {
			return true
		}
	}

	return false
}
