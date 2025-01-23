package model

import (
	"github.com/uptrace/bun"
)

type User struct {
    bun.BaseModel `bun:"table:users"`
    ID            int64  `bun:",pk" json:"id"`
    Name          string `bun:"name,notnull" json:"name"`
    Email         string `bun:"email,unique,notnull" json:"email"`
}
