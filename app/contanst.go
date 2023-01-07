package app

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Status uint

var Context = context.Background()

var DBOp = &sqlx.DB{}

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
