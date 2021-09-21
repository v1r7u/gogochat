package inmemory

import (
	"errors"
	"fmt"
	ggc "gogochat"
)

var maxUserId = 0

type UserService struct {
	db map[int]*ggc.User
}

func NewUserService() *UserService {
	return &UserService{db: make(map[int]*ggc.User)}
}

func (s *UserService) FindUserById(id int) (*ggc.User, error) {
	for _, user := range s.db {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("no user with id %v", id))
}

func (s *UserService) FindUserByUsername(username string) (*ggc.User, error) {
	for _, user := range s.db {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("no user with username %v", username))
}

func (s *UserService) CreateUser(user *ggc.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	for _, u := range s.db {
		if u.Username == user.Username {
			return errors.New(fmt.Sprintf("%v already exists", u.Username))
		}
	}

	maxUserId++
	user.Id = maxUserId
	s.db[maxUserId] = user

	return nil
}

