package db

import (
    "context"
    "errors"
    m "github.com/baisystems/go-test/internal/pkg/model"
)

type MockUserRepository struct {
    users map[int64]*m.User
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{users: make(map[int64]*m.User)}
}

func (u *MockUserRepository) CreateUser(ctx context.Context, user *m.User) error {
    if _, exists := u.users[user.ID]; exists {
        return errors.New("user already exists")
    }
    u.users[user.ID] = user
    return nil
}

func (u *MockUserRepository) GetUserByID(ctx context.Context, id int64) (*m.User, error) {
    user, exists := u.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (u *MockUserRepository) UpdateUser(ctx context.Context, user *m.User) error {
    if _, exists := u.users[user.ID]; !exists {
        return errors.New("user not found")
    }
    u.users[user.ID] = user
    return nil
}

func (u *MockUserRepository) DeleteUser(ctx context.Context, id int64) error {
    if _, exists := u.users[id]; !exists {
        return errors.New("user not found")
    }
    delete(u.users, id)
    return nil
}
