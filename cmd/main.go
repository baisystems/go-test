package main

import (
	"context"
	"fmt"
	"log"
    "github.com/uptrace/bun"

	"github.com/baisystems/go-test/internal/pkg/config"
	"github.com/baisystems/go-test/internal/pkg/database"
	"github.com/baisystems/go-test/internal/pkg/model"
	"github.com/baisystems/go-test/internal/pkg/service"
)

func main() {

	config := &config.Config{
		Ctx: context.Background(),
	}

	userDB := database.NewUserDatabase(config, &bun.DB{})
	userDB.Init()

	// Create a new UserService with the real database
	userService := service.NewUserService(config, userDB)

	// Insert a new user with a primary key 1
	user := &model.User{ID: 1, Name: "Charles Bay", Email: "name@company.com"}
	if err := userService.CreateUser(user); err != nil {
		log.Fatal(err)
	}

	// Select a user by the primary key 1.
	// user := new(model.User)  // for newer go

	user = &model.User{ID: 1}
	retrievedUser, err := userService.GetUser(user.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Retrieved User: %+v\n", retrievedUser)

}
