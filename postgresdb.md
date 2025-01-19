package main

import (
    "context"
    "errors"
)

type UserDatabase struct {
    users map[int64]*User
}

func NewUserDatabase() *UserDatabase {
    return &UserDatabase{users: make(map[int64]*User)}
}

func (m *UserDatabase) CreateUser(ctx context.Context, user *User) error {
    if _, exists := m.users[user.ID]; exists {
        return errors.New("user already exists")
    }
    m.users[user.ID] = user
    return nil
}

func (m *UserDatabase) GetUser(ctx context.Context, id int64) (*User, error) {
    user, exists := m.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (m *UserDatabase) UpdateUser(ctx context.Context, user *User) error {
    if _, exists := m.users[user.ID]; !exists {
        return errors.New("user not found")
    }
    m.users[user.ID] = user
    return nil
}

func (m *UserDatabase) DeleteUser(ctx context.Context, id int64) error {
    if _, exists := m.users[id]; !exists {
        return errors.New("user not found")
    }
    delete(m.users, id)
    return nil
}
