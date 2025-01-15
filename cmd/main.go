package main

import (
    "context"
    "fmt"
    "github.com/uptrace/bun"
    "github.com/uptrace/bun/driver/sqliteshim"
    "github.com/uptrace/bun/dialect/sqlitedialect"
    "github.com/uptrace/bun/extra/bundebug"
    "log"
	m "github.com/baisystems/go-test/internal/pkg/model"
    "githube.com/baisystems/go-test/internal/pkg/db"
)

var db *bun.DB
func main() {
    // Connect to SQLite database
    dsn := "file::memory:?cache=shared"
    sqldb := sqliteshim.Open(dsn)
    defer sqldb.Close()

    db = bun.NewDB(sqldb, sqlitedialect.New())
    db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

    // Create table
    ctx := context.Background()
    if _, err := db.NewCreateTable().Model((*m.User)(nil)).Exec(ctx); err != nil {
        log.Fatal(err)
    }
}

