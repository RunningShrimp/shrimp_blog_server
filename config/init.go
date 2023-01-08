package config

import (
	"fmt"
	redisClient "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"shrimp_blog_sever/framework/app"
	"shrimp_blog_sever/framework/config"
)

/*
这是解析配置文件的，是整个应用的首先开始的地方，这里只要的出错就是严重的问题，所以需要整个应用退出
*/

func init() {
	config.AppInit = &appInit{}
}

type appInit struct {
}

func (a *appInit) InitDBClient(wrapperConfig config.WrapperConfig) {
	database := wrapperConfig.Config.Database
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", database.Username, database.Password, database.URL, database.Dbname)
	db, err := sqlx.Open(database.Type, connectInfo)
	if err != nil {
		panic(err)
	}
	app.DBClient = db
}

func (a *appInit) InitRedisClient(wrapperConfig config.WrapperConfig) {
	redisConfig := wrapperConfig.Config.Redis

	app.RedisClient = redisClient.NewClient(&redisClient.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	})
}
