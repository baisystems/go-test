package user

import (
	"context"
	"fmt"
	m "github.com/baisystems/go-test/internal/pkg/model"
	"githube.com/baisystems/go-test/internal/pkg/db"
	"log"
)

func setupTestUser() error {
	ctx := context.Background()

	// Insert a new user
	user := &m.User{Name: "Charles Bai", Email: "email@example.com"}
	if _, err := db.NewInsert().Model(user).Exec(ctx); err != nil {
		log.Fatal(err)
	}

	// Select user by ID
	var selectedUser m.User
	if err := db.NewSelect().Model(&selectedUser).Where("id = ?", user.ID).Scan(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Selected User: %+v\n", selectedUser)

	// Update user
	user.Name = "Jane Doe"
	if _, err :=db.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(ctx); err != nil {
		log.Fatal(err)
	}

	// Select user by ID
	if err :=db.NewSelect().Model(&selectedUser).Where("id = ?", user.ID).Scan(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Selected User: %+v\n", selectedUser)

	// Delete user
	if _, err :=db.NewDelete().Model(user).Where("id = ?", user.ID).Exec(ctx); err != nil {
		log.Fatal(err)
	}

	// Select user by ID
	if err :=db.NewSelect().Model(&selectedUser).Where("id = ?", user.ID).Scan(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Selected User: %+v\n", selectedUser)

	return nil
}
