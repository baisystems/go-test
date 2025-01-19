package database

import (
	"errors"

	"github.com/baisystems/go-test/internal/pkg/config"
	"github.com/baisystems/go-test/internal/pkg/model"
)

type MockUserDatabase struct {
	Config *config.Config
	Users  map[int64]*model.User
}

func NewMockUserDatabase(config *config.Config, users map[int64]*model.User) *MockUserDatabase {
	return &MockUserDatabase{Config: config, Users: users}
}

func (u *MockUserDatabase) Init() error {
    if u.Users == nil {
        u.Users = make(map[int64]*model.User)
    }
    return nil
}

func (u *MockUserDatabase) CreateUser(user *model.User) error {
	if _, exists := u.Users[user.ID]; exists {
		return errors.New("user already exists")
	}
	u.Users[user.ID] = user
	return nil
}

func (u *MockUserDatabase) GetUser(id int64) (*model.User, error) {
	user, exists := u.Users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *MockUserDatabase) UpdateUser(user *model.User) error {
	if _, exists := u.Users[user.ID]; !exists {
		return errors.New("user not found")
	}
	u.Users[user.ID] = user
	return nil
}

func (u *MockUserDatabase) DeleteUser(id int64) error {
	if _, exists := u.Users[id]; !exists {
		return errors.New("user not found")
	}
	delete(u.Users, id)
	return nil
}
