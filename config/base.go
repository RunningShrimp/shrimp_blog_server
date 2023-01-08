package config

import (
	"fmt"
	"shrimp_blog_sever/framework/config"
)

func ServerPort() string {
	port := config.RsConfig.Config.Server.Port

	if port < 0 || port > 65535 {
		panic("端口不正确")
	}

	return fmt.Sprintf(":%d", port)
}
