package api

import (
	"fmt"
	"shrimp_blog_sever/framework/app"
)

type UserController struct {
	app.BaseController
}
type UserRequest struct {
	app.BaseRequest
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u UserController) Login(request UserRequest) string {
	return fmt.Sprintf("hello, %s,你的密码是：%s", request.Username, request.Password)
}
