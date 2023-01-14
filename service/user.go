package service

import (
	"shrimp_blog_sever/dao"
	"shrimp_blog_sever/framework/app"
)

type UserService struct{}

var userDao dao.UserDao = &dao.UserDaoImpl{}

func (s *UserService) Login(shCtx *app.RsContext) string {
	return "user login"
}
