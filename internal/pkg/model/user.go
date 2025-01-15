package models

import (
	"github.com/uptrace/bun"
)

type User struct { 
	bun.BaseModel `bun:"table:users"` 
	ID int64 `bun:",pk,autoincrement"` 
	Name string `bun:"name,notnull"` 
	Email string `bun:"email,unique,notnull"` 
}