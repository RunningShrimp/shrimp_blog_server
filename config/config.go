package config

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	redisClient "github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	constant "shrimp_blog_sever/app"
	"strings"
)

/*
这是解析配置文件的，是整个应用的首先开始的地方，这里只要的出错就是严重的问题，所以需要整个应用退出
*/

var (
	RsConfig = &AppConfig{
		EnvCfgFilePath: "",
		Env:            "",
		WrapperConfig:  WrapperConfig{},
	}
)

func Init(path string) {
	initConfig(path)
	initDatabase()
	initRedis()
}

func getEnvCfgFilePath(EntranceFile string) string {
	resourceDir, resourceFile := getEntranceDirectory(EntranceFile)

	envType := getEnvType(resourceFile)

	err := filepath.WalkDir(
		resourceDir,
		func(path string, d fs.DirEntry, err error) error {
			fileInfo, err := d.Info()
			if err != nil {
				return err
			}

			// todo: 多个文件怎么办？
			if strings.Contains(fileInfo.Name(), envType) {
				RsConfig.EnvCfgFilePath = path
			}

			return err
		},
	)
	if err != nil {
		panic(err)
	}

	return RsConfig.EnvCfgFilePath
}
func getEnvType(appFile string) string {
	type (
		appInfo struct {
			Profile struct {
				Active string `yaml:"active"`
			} `yaml:"profile"`
		}
	)

	pro := appInfo{}

	content, err := os.ReadFile(appFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &pro)
	if err != nil {
		panic(err)
	}

	RsConfig.Env = pro.Profile.Active

	return pro.Profile.Active
}

func getEntranceDirectory(configPathPrefix string) (string, string) {
	appEntranceDirectory, err := filepath.Abs(configPathPrefix)
	if err != nil {
		panic(err)
	}

	return appEntranceDirectory, filepath.Join(appEntranceDirectory, "resource", "app.yaml")
}

func initConfig(configFile string) {
	configContent, err := os.ReadFile(getEnvCfgFilePath(configFile))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configContent, &RsConfig.WrapperConfig)
	if err != nil {
		panic(err)
	}
}
func initDatabase() {
	database := RsConfig.Config.Database
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", database.Username, database.Password, database.URL, database.Dbname)
	db, err := sqlx.Open(database.Type, connectInfo)
	if err != nil {
		panic(err)
	}
	constant.DBClient = db
}

func initRedis() {
	redisConfig := RsConfig.Config.Redis

	constant.RedisClient = redisClient.NewClient(&redisClient.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	})
}

func ServerPort() string {
	//todo:check port valid
	return fmt.Sprintf(":%d", RsConfig.Config.Server.Port)
}
