package main

import (
	_ "shrimp_blog_sever/config"
	"shrimp_blog_sever/framework/config"
	"shrimp_blog_sever/framework/server"
	_ "shrimp_blog_sever/routes"
)

func init() {
	config.Init(".")
}

func main() {
	server.NewServer().Run()
}
