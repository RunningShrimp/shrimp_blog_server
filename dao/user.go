package dao

import (
	"fmt"
	"go.uber.org/zap"
	"shrimp_blog_sever_v2/constant"
	"shrimp_blog_sever_v2/log"
	"shrimp_blog_sever_v2/model"
	"shrimp_blog_sever_v2/utils"
)

type UserDao BaseDao[model.User]

type UserDaoImpl struct{}

func (u *UserDaoImpl) SelectById(id int) *model.User {
	sql := fmt.Sprintf("select %s from user where id = %d ", GenFiledString[model.User](), id)

	log.Logger.Info("sql语句：", zap.String("sql", sql))
	slice := rowsToSlice[model.User](sql)
	if slice != nil {
		return slice[0]
	}
	return nil
}

func (u *UserDaoImpl) SelectByPage(page, pageSize int) []*model.User {
	sql := fmt.Sprintf("select %s from user limit %d , %d", GenFiledString[model.User](), (page-1)*pageSize, pageSize)
	log.Logger.Info("sql语句：", zap.String("sql", sql))
	return rowsToSlice[model.User](sql)

}

func (u *UserDaoImpl) SelectByIds(ids ...int) []*model.User {
	sql := fmt.Sprintf("select %s from user where id in ( %s )", GenFiledString[model.User](), utils.JoinIntSlice(ids, ","))
	log.Logger.Info("sql语句：", zap.String("sql", sql))
	return rowsToSlice[model.User](sql)
}

func (u *UserDaoImpl) SelectStatusById(id int) int {
	var status int
	sql := fmt.Sprintf("select status from user where  id = %d", id)

	if err := constant.DBOp.Get(&status, sql); err != nil {
		log.Logger.Error("查询出错：", zap.Error(err))
	}

	return status
}

func (u *UserDaoImpl) Insert(t *model.User) {
	constant.DBOp.Exec()
}

func (u *UserDaoImpl) Update(t *model.User) {
	//TODO implement me
	panic("implement me")
}

func (u *UserDaoImpl) DeleteById(t *model.User) {
	//TODO implement me
	panic("implement me")
}
