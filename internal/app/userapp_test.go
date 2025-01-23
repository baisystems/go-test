package app

import (
    "testing"
	"context"

	"github.com/baisystems/go-test/internal/pkg/config"
	"github.com/baisystems/go-test/internal/pkg/database"
	"github.com/baisystems/go-test/internal/pkg/model"
	"github.com/baisystems/go-test/internal/pkg/service"
)

func TestUserService_CreateUser(t *testing.T) {
	config := &config.Config{
		Ctx: context.Background(),
	}
	mockDB := database.NewMockUserDatabase(config, make(map[int64]*model.User))
    userService := service.NewUserService(config, mockDB)
	
    user := &model.User{ID: 2, Name: "Bob", Email: "bob@example.com"}
    err := userService.CreateUser(user)
    if err != nil {
		t.Fatalf("expected no error, got %v", err)
    }
	
    createdUser, err := userService.GetUser(2)
    if err != nil {
		t.Fatalf("expected no error, got %v", err)
    }
    if createdUser.Name != "Bob" {
		t.Fatalf("expected name to be Bob, got %v", createdUser.Name)
    }
}

func TestUserService_GetUser(t *testing.T) {
	users := make(map[int64]*model.User)
	users[1] = &model.User{
		ID: 1, Name: "Alice", Email: "alice@example.com",
	}

	config := &config.Config{
		Ctx: context.Background(),
	}
	mockDB := database.NewMockUserDatabase(config, users)
    userService := service.NewUserService(config, mockDB)


	user, err := userService.GetUser(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.Name != "Alice" {
		t.Fatalf("expected name to be Alice, got %v", user.Name)
	}
}