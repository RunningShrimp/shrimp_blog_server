package model

// LinkDao is the manager for logic model data accessing and custom defined data operations functions management.
type LinkDao struct {
	Table   string      // Table is the underlying table name of the DAO.
	Group   string      // Group is the database configuration group name of current DAO.
	Columns LinkColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// LinkColumns defines and stores column names for table link.
type LinkColumns struct {
	ID         string //
	Name       string // 链接名
	URL        string // 链接地址
	Avatar     string // 链接图标
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// linkColumns holds the columns for table link.
var linkColumns = LinkColumns{
	ID:         "id",
	Name:       "name",
	URL:        "url",
	Avatar:     "avatar",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewLinkDao creates and returns a new DAO object for table data access.
func NewLinkDao() *LinkDao {
	return &LinkDao{
		Group:   "default",
		Table:   "link",
		Columns: linkColumns,
	}
}
