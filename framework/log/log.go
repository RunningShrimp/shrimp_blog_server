package log

import (
	"go.uber.org/zap"
	"shrimp_blog_sever/framework/config"
)

var Logger *zap.Logger

func init() {

	var err error

	switch config.RsConfig.Env {
	case "dev":
		Logger, err = zap.NewDevelopment()
	case "prd":
		Logger, err = zap.NewProduction()
	default:
		Logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}
}
