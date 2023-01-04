package dao

import (
	"go.uber.org/zap"
	"reflect"
	"shrimp_blog_sever_v2/constant"
	"shrimp_blog_sever_v2/log"
	"shrimp_blog_sever_v2/model"
	"strings"
	"time"
)

type Time time.Time

// Model 类型集合
type Model interface {
	model.User | model.Article | model.Comment | model.ArticleCategory | model.ArticleLabel | model.Category | model.Label |
		model.Link | model.Messages | model.Navbar | model.PageTemplate | model.Resource | model.Role | model.RoleResources | model.Subject | model.SubjectItem | model.UserRole
}

type Table interface {
	model.User | model.Article | model.Comment | model.ArticleCategory | model.ArticleLabel | model.Category | model.Label |
		model.Link | model.Messages | model.Navbar | model.PageTemplate | model.Resource | model.Role | model.RoleResources | model.Subject | model.SubjectItem | model.UserRole
}

type BaseDao[T Model] interface {
	SelectById(id int) *T
	SelectByPage(page, pageSize int) []*T
	SelectByIds(ids ...int) []*T
	SelectStatusById(id int) int
	Insert(t *T)
	Update(t *T)
	DeleteById(t *T)
}

const SqlSelectById = "select "

// 生成select 查询字段
func genSelectFiledString[T Model]() string {
	var m T
	var result = make([]string, 0)

	st := reflect.TypeOf(m)
	fieldNum := st.NumField()
	for i := 0; i < fieldNum; i++ {
		dbTag := st.Field(i).Tag.Get("db")
		result = append(result, dbTag)
	}

	return strings.Join(result, ",")
}

func genInsertFiledString[T Model]() string {
	var m T
	var result = make([]string, 0)

	st := reflect.TypeOf(m)
	fieldNum := st.NumField()
	for i := 0; i < fieldNum; i++ {
		dbTag := st.Field(i).Tag.Get("db")
		if strings.Contains(dbTag, "id") || strings.Contains(strings.ToLower(dbTag), "time") {
			continue
		}
		result = append(result, dbTag)
	}

	return strings.Join(result, ",")
}

func genInsertValueString[T Model](m T) string {
	var result = make([]any, 0)

	st := reflect.TypeOf(m)
	fieldNum := st.NumField()
	for i := 0; i < fieldNum; i++ {
		v := reflect.ValueOf(m)
	}
}

func rowsToSlice[T Model](sql string) []*T {
	var models []*T

	stmt, err := constant.DBOp.Prepare(sql)
	if err != nil {
		log.Logger.Error("sql语句错误：", zap.Error(err))
		return nil
	}
	stmt.QueryRow()

	rows, err := constant.DBOp.Queryx(sql)
	if err != nil {
		log.Logger.Error("查询错误，", zap.Error(err))
		return nil
	}

	for rows.Next() {
		var t T
		if err := rows.StructScan(&t); err != nil {
			log.Logger.Error("映射错误：", zap.Error(err))
		}
		models = append(models, &t)
	}

	return models
}
