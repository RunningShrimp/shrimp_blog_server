package router

import "shrimp_blog_sever/framework/app"

type node struct {
	path       string
	children   []*node
	controller app.IController
	function   any
	method     string
}
