package routes

import (
	"shrimp_blog_sever/api"
	"shrimp_blog_sever/framework/server"
)

func init() {
	//router.AddRouter("/user/login", &api.UserController{},)\
	server.Post("/user/login", api.UserController{}.Login)
	server.Get("/user/info", api.UserController{}.Info)
}
