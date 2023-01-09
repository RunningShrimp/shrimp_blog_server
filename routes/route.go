package routes

import (
	"shrimp_blog_sever/api"
	"shrimp_blog_sever/framework/router"
)

func init() {
	router.AddRouter("/user/login", &api.UserController{})
}
