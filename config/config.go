// Package config 这是解析配置文件的，是整个应用的首先开始的地方，这里只要的出错就是严重的问题，所以需要整个应用退出
package config

import (
	"context"
	"fmt"
	redisClient "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	"shrimp_blog_sever_v2/config/internal"
	"shrimp_blog_sever_v2/constant"
	"strings"
)

func InitGoble(path string) {
	parse(path)
	newDatabase()
	newRedis()
}

var app = &internal.App{}

func path(path string) string {
	resourcePath := resourcePath(path)

	env := activeEnv(path)

	file := filepath.Join(resourcePath, "app.yaml")

	filepath.WalkDir(
		resourcePath,
		func(path string, d fs.DirEntry, err error) error {
			fileInfo, err := d.Info()
			if err != nil {
				return err
			}
			if strings.Contains(fileInfo.Name(), env) {
				file = path
			}
			return err
		},
	)

	// 放入context中，方便以后取值
	constant.Context = context.WithValue(constant.Context, constant.ConfigPath, file)
	return file
}

func activeEnv(path string) string {
	appFilePath := filepath.Join(resourcePath(path), "app.yaml")

	type profile struct {
		Active string `yaml:"active"`
	}
	type appInfo struct {
		Profile profile `yaml:"profile"`
	}

	pro := appInfo{}

	content, err := os.ReadFile(appFilePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &pro)
	if err != nil {
		panic(err)
	}

	constant.Context = context.WithValue(
		constant.Context,
		constant.EnvConfig,
		pro.Profile.Active,
	)

	return pro.Profile.Active
}

func resourcePath(path string) string {
	currentProgramPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return filepath.Join(currentProgramPath, "resource")
}

func parse(configFile string) {
	content, err := os.ReadFile(path(configFile))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &app)
	if err != nil {
		panic(err)
	}
}
func newDatabase() {
	database := app.Config.Database
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", database.Username, database.Password, database.URL, database.Dbname)
	db, err := sqlx.Open(database.Type, connectInfo)
	if err != nil {
		panic(err)
	}
	constant.DBOp = db
}

func newRedis() {
	redisConfig := app.Config.Redis

	constant.RedisClient = redisClient.NewClient(&redisClient.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	})
}
