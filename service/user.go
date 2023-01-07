package service

import (
	"shrimp_blog_sever/app"
	"shrimp_blog_sever/dao"
)

type UserService struct{}

var userDao dao.UserDao = &dao.UserDaoImpl{}

func (s *UserService) Login(shCtx *app.RsContext) string {
	return "user login"
}
