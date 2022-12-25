package model

// SubjectItemDao is the manager for logic model data accessing and custom defined data operations functions management.
type SubjectItemDao struct {
	Table   string             // Table is the underlying table name of the DAO.
	Group   string             // Group is the database configuration group name of current DAO.
	Columns SubjectItemColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SubjectItemColumns defines and stores column names for table subject_item.
type SubjectItemColumns struct {
	ID         string //
	SubjectID  string // 专题uid
	ArticleID  string // 文章uid
	Status     string // 状态
	Sort       string // 排序字段
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
}

// subjectItemColumns holds the columns for table subject_item.
var subjectItemColumns = SubjectItemColumns{
	ID:         "id",
	SubjectID:  "subject_id",
	ArticleID:  "article_id",
	Status:     "status",
	Sort:       "sort",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
}

// NewSubjectItemDao creates and returns a new DAO object for table data access.
func NewSubjectItemDao() *SubjectItemDao {
	return &SubjectItemDao{
		Group:   "default",
		Table:   "subject_item",
		Columns: subjectItemColumns,
	}
}
