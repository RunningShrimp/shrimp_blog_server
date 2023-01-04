package internal

// NavbarTable is the manager for logic model data accessing and custom defined data operations functions management.
type NavbarTable struct {
	Table string // Table is the underlying table name of the Table.

	Columns NavbarColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// NavbarColumns defines and stores column names for table navbar.
type NavbarColumns struct {
	ID          string //
	Name        string // 导航名
	URL         string // 链接地址
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	DeleteTime  string // 删除时间
	Status      string // 状态
	NavbarLevel string // 导航级别
	ParentID    string // 父级id
	Sort        string // 排序
}

// navbarColumns holds the columns for table navbar.
var navbarColumns = NavbarColumns{
	ID:          "id",
	Name:        "name",
	URL:         "url",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	DeleteTime:  "delete_time",
	Status:      "status",
	NavbarLevel: "navbar_level",
	ParentID:    "parent_id",
	Sort:        "sort",
}

// NewNavbarTable creates and returns a new Table object for table data access.
func NewNavbarTable() *NavbarTable {
	return &NavbarTable{

		Table:   "navbar",
		Columns: navbarColumns,
	}
}
