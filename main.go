package main

import (
	"net/http"
	"shrimp_blog_sever/config"
)

func init() {
	config.Init(".")
}
func main() {

	//serveMux := http.NewServeMux()

	err := http.ListenAndServe(config.ServerPort(), nil)
	if err != nil {
		panic(err)
	}
}
