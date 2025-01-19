package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"

	"github.com/baisystems/go-test/internal/pkg/config"
	"github.com/baisystems/go-test/internal/pkg/model"
)

type UserDatabase struct {
	Config 	*config.Config
	Db		*bun.DB
}

func NewUserDatabase(config *config.Config, db *bun.DB) *UserDatabase {
	return &UserDatabase{Config: config, Db: db}
}

func (u *UserDatabase) Init() error {

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Traverse up the directory tree until you find the project root
	root := cwd
	for {
		if _, err := os.Stat(path.Join(root, "go.mod")); err == nil {
			break
		}
		parent := path.Dir(root)
		if parent == root {
			fmt.Println("Error: Project root not found")
			return err
		}
		root = parent
	}

	// Connect to SQLite database. If use as in memory:  dsn := "file::memory:?cache=shared"
	dsn := fmt.Sprintf("file:%s/data/sqlite.db", root)
	fmt.Println("dsn string: ", dsn)
	sqldb, err := sql.Open(sqliteshim.ShimName, dsn)
	if err != nil {
		panic(err)
	}
	// If you are using an in-memory database, you need to configure *sql.DB to NOT close active connections.
	// sqldb.SetMaxIdleConns(1000)
	// sqldb.SetConnMaxLifetime(0)

	db := bun.NewDB(sqldb, sqlitedialect.New())
	// db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	u.Db = db

	// Check if the table exists
	var exists bool
	query := `SELECT count(*) FROM sqlite_master WHERE type='table' AND name='users';`
	if err := u.Db.NewRaw(query).Scan(u.Config.Ctx, &exists); err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Println("Table exists")
	} else {
		// Create users table.
		if _, err := u.Db.NewCreateTable().Model((*model.User)(nil)).Exec(u.Config.Ctx); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func (u *UserDatabase) CreateUser(user *model.User) error {
	fmt.Printf("Creating user: %+v to database \n", user)

	existingUser, err := u.GetUser(user.ID)
	if err != nil {
		// not existing, do insert
		res, err := u.Db.NewInsert().Model(user).Exec(u.Config.Ctx)
		if err != nil {
			log.Fatal(err)
			return err
		}

		n, _ := res.RowsAffected()
		fmt.Printf("insert succeed, rows affected: %d \n", n)

		return nil

	}

	if equal := compareUsers(*existingUser, *user); !equal {
		return u.UpdateUser(user)
	}
	fmt.Println("user existing and no change on data, do nothing")
	return nil
}

func (u *UserDatabase) GetUser(id int64) (*model.User, error) {
	fmt.Printf("GetUser with ID: %d \n", id)

	user := &model.User{ID: id, Name: "", Email: ""}
	if err := u.Db.NewSelect().Model(user).Where("id = ?", user.ID).Scan(u.Config.Ctx); err != nil {
		log.Fatal(err)
		return &model.User{}, err
	}
	return user, nil
}

func (u *UserDatabase) UpdateUser(user *model.User) error {
	// if _, exists := u.Users[user.ID]; !exists {
	//     return errors.New("user not found")
	// }
	// u.Users[user.ID] = user

	res, err := u.Db.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(u.Config.Ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}

	n, _ := res.RowsAffected()
	fmt.Printf("update succeed, rows affected: %d \n", n)

	return nil
}

func (u *UserDatabase) DeleteUser(id int64) error {
	// if _, exists := u.Users[id]; !exists {
	// 	return errors.New("user not found")
	// }
	// delete(u.Users, id)
	return nil
}

func compareUsers(u1 model.User, u2 model.User) bool {
	return u1.ID == u2.ID && u1.Name == u2.Name && u1.Email == u2.Email
}
