package service

import (
	"github.com/baisystems/go-test/internal/pkg/config"
	"github.com/baisystems/go-test/internal/pkg/database"
	"github.com/baisystems/go-test/internal/pkg/model"
)

type UserService struct {
	Config *config.Config
	Db     database.UserDatabaseInterface
}

func NewUserService(config *config.Config, db database.UserDatabaseInterface) *UserService {
	return &UserService{Config: config, Db: db}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Db.CreateUser(user)
}

func (s *UserService) GetUser(id int64) (*model.User, error) {
	return s.Db.GetUser(id)
}
