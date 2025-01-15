package db

import (
    "context"
    m "github.com/baisystems/go-test/internal/pkg/model"
)

type UserRepositoryInterface interface {
    CreateUser(ctx context.Context, user *m.User) error
    GetUserByID(ctx context.Context, id int64) (*m.User, error)
    UpdateUser(ctx context.Context, user *m.User) error
    DeleteUser(ctx context.Context, id int64) error
}

