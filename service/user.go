package service

import "shrimp_blog_sever_v2/dao"

type UserService struct{}

var userDao dao.UserDao = &dao.UserDaoImpl{}

func (s *UserService) Login() string {
	return "user login"
}
