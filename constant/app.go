package constant

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"os"
)

type ContextKey int

var Separator = string(os.PathSeparator)
var Context = context.Background()

var DBOp = &sqlx.DB{}

var RedisClient = &redis.Client{}

const (
	ConfigPath ContextKey = iota
	EnvConfig
	DbConfig
)
