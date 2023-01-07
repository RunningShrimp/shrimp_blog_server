package log

import (
	"go.uber.org/zap"
	"shrimp_blog_sever_v2/app"
	"strings"
)

var Logger *zap.Logger

func init() {
	env := "dev"
	if v := app.Context.Value(app.EnvConfig); v != nil {
		env = v.(string)
	}

	var err error
	if strings.Contains(env, "prd") {
		Logger, err = zap.NewProduction()
	} else {
		Logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
}
