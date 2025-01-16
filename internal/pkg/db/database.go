package db

import (
    m "github.com/baisystems/go-test/internal/pkg/model"
    "github.com/baisystems/go-test/internal/pkg/config"
)

type UserRepositoryInterface interface {
    CreateUser(config *config.Config, user *m.User) error
    GetUserByID(config *config.Config, id int64) (*m.User, error)
    UpdateUser(config *config.Config, user *m.User) error
    DeleteUser(config *config.Config, id int64) error
}

