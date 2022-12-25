package model

// ResourceDao is the manager for logic model data accessing and custom defined data operations functions management.
type ResourceDao struct {
	Table   string          // Table is the underlying table name of the DAO.
	Group   string          // Group is the database configuration group name of current DAO.
	Columns ResourceColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// ResourceColumns defines and stores column names for table resource.
type ResourceColumns struct {
	ID         string //
	Name       string // 资源名
	URL        string // 资源路径
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// resourceColumns holds the columns for table resource.
var resourceColumns = ResourceColumns{
	ID:         "id",
	Name:       "name",
	URL:        "url",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewResourceDao creates and returns a new DAO object for table data access.
func NewResourceDao() *ResourceDao {
	return &ResourceDao{
		Group:   "default",
		Table:   "resource",
		Columns: resourceColumns,
	}
}