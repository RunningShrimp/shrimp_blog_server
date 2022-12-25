package model

// SubjectDao is the manager for logic model data accessing and custom defined data operations functions management.
type SubjectDao struct {
	Table   string         // Table is the underlying table name of the DAO.
	Group   string         // Group is the database configuration group name of current DAO.
	Columns SubjectColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SubjectColumns defines and stores column names for table subject.
type SubjectColumns struct {
	ID         string //
	Name       string // 专题名
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// subjectColumns holds the columns for table subject.
var subjectColumns = SubjectColumns{
	ID:         "id",
	Name:       "name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewSubjectDao creates and returns a new DAO object for table data access.
func NewSubjectDao() *SubjectDao {
	return &SubjectDao{
		Group:   "default",
		Table:   "subject",
		Columns: subjectColumns,
	}
}
