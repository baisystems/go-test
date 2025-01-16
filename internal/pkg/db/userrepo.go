package db

import (
    "fmt"
    "log"
    "errors"
	m "github.com/baisystems/go-test/internal/pkg/model"
    "github.com/baisystems/go-test/internal/pkg/config"
)

type UserRepository struct {
    users map[int64]*m.User
}

func NewUserRepository() *UserRepository {
    return &UserRepository{users: make(map[int64]*m.User)}
}

func (u *UserRepository) CreateUser(config *config.Config, user *m.User) error {
    fmt.Printf("Creating user: %+v to database \n", user)

    // if _, exists := u.users[user.ID]; exists {
    //     return errors.New("user already exists")
    // }

    // Insert a new user with a primary key
	res, err := config.Db.NewInsert().Model(user).Exec(config.Ctx)
    if err != nil {
        log.Fatal(err)
		return err
    }

    n, _ := res.RowsAffected()
	fmt.Printf("rows affected: %d \n", n)

    u.users[user.ID] = user
    return nil
}

// func (u *UserRepository) GetUserByIDFromCache(config *config.Config, id int64) (m.User, error) {
//     user, exists := u.users[id]
//     if !exists {
//         return m.User{}, errors.New("user not found")
//     }
//     return *user, nil
// }

func (u *UserRepository) GetUserByID(config *config.Config, id int64) (m.User, error) {
    fmt.Printf("GetUserByID with ID: %d \n", id)

    user := &m.User{ID: id, Name: "", Email: ""}
	if err := config.Db.NewSelect().Model(user).Where("id = ?", user.ID).Scan(config.Ctx); err != nil {
		log.Fatal(err)
		return m.User{}, err
	}
	return *user, nil
}

func (u *UserRepository) UpdateUser(config *config.Config, user *m.User) error {
    if _, exists := u.users[user.ID]; !exists {
        return errors.New("user not found")
    }
    u.users[user.ID] = user
    return nil
}

func (u *UserRepository) DeleteUser(config *config.Config, id int64) error {
    if _, exists := u.users[id]; !exists {
        return errors.New("user not found")
    }
    delete(u.users, id)
    return nil
}
