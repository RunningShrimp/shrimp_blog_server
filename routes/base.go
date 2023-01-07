package routes

import (
	"net/http"
	"shrimp_blog_sever_v2/app"
)

type router struct { //nolint:maligned
	Path       string
	Method     string
	Controller app.BaseController
}

var Routers = make([]*router, 8)

func (r *router) AddRouter(path, method string, controller app.BaseController) {
	Routers = append(Routers, &router{
		Path:       path,
		Method:     method,
		Controller: controller,
	})
	switch method {
	case http.MethodPost:

	case http.MethodGet:

	}

}

func (r *router) Get(path string, controller app.BaseController) {
	r.AddRouter(path, http.MethodGet, controller)

}
func (r *router) Post(path string, controller app.BaseController) {
	r.AddRouter(path, http.MethodPost, controller)
}
