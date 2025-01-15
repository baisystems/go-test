package db

import (
    "log"
    "context"
    "errors"
	m "github.com/baisystems/go-test/internal/pkg/model"
)

type UserRepository struct {
    users map[int64]*m.User
}

func NewUserRepository() *UserRepository {
    return &UserRepository{users: make(map[int64]*m.User)}
}

func (u *UserRepository) CreateUser(cfg config, user *m.User) error {
    if _, exists := u.users[user.ID]; exists {
        return errors.New("user already exists")
    }
	if _, err := db.NewInsert().Model(user).Exec(ctx); err != nil {
        log.Fatal(err)
		return err
    }
    u.users[user.ID] = user
    return nil
}

func (u *UserRepository) GetUserByID(ctx context.Context, id int64) (m.User, error) {
    user, exists := u.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, user *m.User) error {
    if _, exists := u.users[user.ID]; !exists {
        return errors.New("user not found")
    }
    u.users[user.ID] = user
    return nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, id int64) error {
    if _, exists := u.users[id]; !exists {
        return errors.New("user not found")
    }
    delete(u.users, id)
    return nil
}
