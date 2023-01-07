package main

import (
	"net/http"
	"shrimp_blog_sever_v2/config"
)

func init() {
	config.InitGoble(".")
}
func main() {

	serveMux := http.NewServeMux()

	err := http.ListenAndServe(config.ServerPort(), nil)
	if err != nil {
		panic(err)
	}
}
