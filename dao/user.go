package dao

import (
	"fmt"
	"go.uber.org/zap"
	"shrimp_blog_sever/app"
	"shrimp_blog_sever/log"
	"shrimp_blog_sever/model"
	"shrimp_blog_sever/utils"
	"time"
)

type UserDao BaseDao[model.User]

type UserDaoImpl struct{}

func (u *UserDaoImpl) SelectById(id int) *model.User {
	sql := fmt.Sprintf("select %s from user where id = %d ", genSelectFiledString[model.User](), id)

	log.Logger.Info("sql语句：", zap.String("sql", sql))
	slice := rowsToSlice[model.User](sql)
	if slice != nil {
		return slice[0]
	}
	return nil
}

func (u *UserDaoImpl) SelectByPage(page, pageSize int) []*model.User {
	sql := fmt.Sprintf("select %s from user limit %d , %d", genSelectFiledString[model.User](), (page-1)*pageSize, pageSize)
	log.Logger.Info("sql语句：", zap.String("sql", sql))
	return rowsToSlice[model.User](sql)
}

func (u *UserDaoImpl) SelectByIds(ids ...int) []*model.User {
	sql := fmt.Sprintf("select %s from user where id in ( %s )", genSelectFiledString[model.User](), utils.JoinIntSlice(",", ids))
	log.Logger.Info("sql语句：", zap.String("sql", sql))
	return rowsToSlice[model.User](sql)
}

func (u *UserDaoImpl) SelectStatusById(id int) int {
	var status int
	sql := fmt.Sprintf("select status from user where  id = %d", id)

	if err := app.DBClient.Get(&status, sql); err != nil {
		log.Logger.Error("查询用户出错：", zap.Error(err))
		return -1
	}

	return status
}

func (u *UserDaoImpl) Insert(t *model.User) bool {
	v, p := genInsertValueAndPlaceholder[model.User](*t)
	sql := fmt.Sprintf("insert into user  %s  VALUES %s ", genInsertFieldString[model.User](), p)
	log.Logger.Info("sql语句：", zap.String("sql", sql))

	result := app.DBClient.MustExec(sql, v...)

	n, err := result.RowsAffected()
	if err != nil {
		log.Logger.Error("插入数据数据库返回错误：", zap.Error(err))
		return false
	}

	return n > 0
}

func (u *UserDaoImpl) Update(t *model.User) bool {
	now := time.Now()
	t.UpdateTime = &now
	sql := fmt.Sprintf("update user set %s where id=:id", genUpdateFieldString[model.User](*t))

	result, err := app.DBClient.NamedExec(sql, *t)
	if err != nil {
		log.Logger.Error("更新用户失败：", zap.Error(err))
		return false
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Logger.Error("数据库返回错误：", zap.Error(err))
		return false
	}

	return affected > 0
}

func (u *UserDaoImpl) DeleteById(t *model.User) bool {
	now := time.Now()
	t.DeleteTime = &now
	t.UpdateTime = &now
	t.Status = uint(app.Disable)

	result, err := app.DBClient.NamedExec("update user set delete_time = :delete_time,status = :status where id = :id", *t)
	if err != nil {
		log.Logger.Error("删除用户失败：", zap.Error(err))
		return false
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Logger.Error("数据库返回错误：", zap.Error(err))
		return false
	}

	return affected > 0
}
