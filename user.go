package gogochat

import (
	"errors"
	"fmt"
)

type User struct {
	Id int `json:"id"`

	Username  string `json:"username"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	return nil
}

func (u User) String() string {
	return fmt.Sprintf("%v user", u.Username)
}

type UserService interface {
	FindUserById(id int) (*User, error)
	FindUserByUsername(username string) (*User, error)
	CreateUser(user *User) error
}
