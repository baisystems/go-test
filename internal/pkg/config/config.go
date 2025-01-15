package config

import (
	"context"

	"github.com/uptrace/bun"
)

type Config struct {
	Ctx context.Context
	Db	*bun.DB
}