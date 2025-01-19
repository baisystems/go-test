package database

import (
	"github.com/baisystems/go-test/internal/pkg/model"
)

type UserDatabaseInterface interface {
	Init() error
	CreateUser(user *model.User) error
	GetUser(id int64) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}
