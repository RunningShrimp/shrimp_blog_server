package main

import (
	"net/http"
	appConfig "shrimp_blog_sever/config"
	"shrimp_blog_sever/framework/config"
)

func init() {
	config.Init(".")
}
func main() {

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/*", func(writer http.ResponseWriter, request *http.Request) {

	})

	err := http.ListenAndServe(appConfig.ServerPort(), serveMux)
	if err != nil {
		panic(err)
	}
}
