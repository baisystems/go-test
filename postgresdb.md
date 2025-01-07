package main

import (
    "context"
    "errors"
)

type UserRepository struct {
    users map[int64]*User
}

func NewUserRepository() *UserRepository {
    return &UserRepository{users: make(map[int64]*User)}
}

func (m *UserRepository) CreateUser(ctx context.Context, user *User) error {
    if _, exists := m.users[user.ID]; exists {
        return errors.New("user already exists")
    }
    m.users[user.ID] = user
    return nil
}

func (m *UserRepository) GetUserByID(ctx context.Context, id int64) (*User, error) {
    user, exists := m.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (m *UserRepository) UpdateUser(ctx context.Context, user *User) error {
    if _, exists := m.users[user.ID]; !exists {
        return errors.New("user not found")
    }
    m.users[user.ID] = user
    return nil
}

func (m *UserRepository) DeleteUser(ctx context.Context, id int64) error {
    if _, exists := m.users[id]; !exists {
        return errors.New("user not found")
    }
    delete(m.users, id)
    return nil
}
