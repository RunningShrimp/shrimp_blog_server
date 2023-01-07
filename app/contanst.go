package app

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Status uint

// request -> mid with RsCtx -> handler(ctx, req, res)
type RsContext struct {
	ctx context.Context
	env string // init
}

func (c RsContext) GetEnv() string {
	return c.GetEnv()
}

func (c RsContext) Context() context.Context {
	return c.ctx
}

var DBClient = &sqlx.DB{}

var RedisClient = &redis.Client{}

const (
	ConfigPath = "configPath"
	EnvConfig  = "env"
	DbConfig   = "db"
)

const (
	Enable  Status = iota + 1
	Disable Status = iota + 1
)
