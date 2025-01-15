package main

import (
	"context"
	"database/sql"
	"fmt"

	// "fmt"
	// "log"
	"github.com/baisystems/go-test/internal/app"
	"github.com/baisystems/go-test/internal/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func main() {
	// Connect to SQLite database
	dsn := "file::memory:?cache=shared"
	sqldb, err := sql.Open(sqliteshim.ShimName, dsn)
	if err != nil {
		panic(err)
	}

    // If you are using an in-memory database, you need to configure *sql.DB to NOT close active connections.
	sqldb.SetMaxIdleConns(1000)
	sqldb.SetConnMaxLifetime(0)

	db := bun.NewDB(sqldb, sqlitedialect.New())
    // db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
    ctx := context.Background()

    config := &config.Config{
        Ctx: ctx,
        Db: db,
    }

    err = user.RunUser(config)
    if err != nil {
        fmt.Println(err.Error())
    }
}
