package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var (
	RsConfig = &AppConfig{
		EnvCfgFilePath: "",
		Env:            "",
		WrapperConfig:  WrapperConfig{},
	}
	AppInit IInit
)

type IInit interface {
	InitDBClient(wrapperConfig WrapperConfig)
	InitRedisClient(wrapperConfig WrapperConfig)
}

func Init(path string) {
	initConfig(path)
	AppInit.InitDBClient(RsConfig.WrapperConfig)
	AppInit.InitRedisClient(RsConfig.WrapperConfig)
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
