package user

import (
	"fmt"
	"github.com/baisystems/go-test/internal/pkg/config"
	m "github.com/baisystems/go-test/internal/pkg/model"
	"log"
)

func RunUser(config *config.Config) error {

	// Create users table.
	if _, err := config.Db.NewCreateTable().Model((*m.User)(nil)).Exec(config.Ctx); err != nil {
		log.Fatal(err)
		return err
	}

	// Insert a new user with a primary key 1
	user := &m.User{ID: 1, Name: "Charles Bai", Email: "name@company.com"}
	res, err := config.Db.NewInsert().Model(user).Exec(config.Ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	n, err := res.RowsAffected()
	fmt.Printf("rows affected: %d \n", n)

	// Select a user by the primary key 1.
	// user := new(m.User)  // for newer go
	selectedUser := &m.User{ID: 1, Name: "", Email: ""}
	if err := config.Db.NewSelect().Model(selectedUser).Where("id = ?", user.ID).Scan(config.Ctx); err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Selected User: %+v\n", selectedUser)

	// // Update user
	// user.Name = "Jane Doe"
	// if _, err := config.Db.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(config.Ctx); err != nil {
	// 	log.Fatal(err)
	// }

	// // Select user by ID
	// if err := config.Db.NewSelect().Model(&selectedUser).Where("id = ?", user.ID).Scan(config.Ctx); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Selected User: %+v\n", selectedUser)

	// // Delete user
	// if _, err := config.Db.NewDelete().Model(user).Where("id = ?", user.ID).Exec(config.Ctx); err != nil {
	// 	log.Fatal(err)
	// }

	// // Select user by ID
	// if err := config.Db.NewSelect().Model(&selectedUser).Where("id = ?", user.ID).Scan(config.Ctx); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Selected User: %+v\n", selectedUser)

	return nil
}
