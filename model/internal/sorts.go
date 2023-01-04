package internal

// CategoryTable is the manager for logic model data accessing and custom defined data operations functions management.
type CategoryTable struct {
	Table string // Table is the underlying table name of the Table.

	Columns CategoryColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// CategoryColumns defines and stores column names for table sorts.
type CategoryColumns struct {
	ID         string //
	Name       string // 分类名
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// sortsColumns holds the columns for table sorts.
var sortsColumns = CategoryColumns{
	ID:         "id",
	Name:       "name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewCategoryTable creates and returns a new Table object for table data access.
func NewCategoryTable() *CategoryTable {
	return &CategoryTable{

		Table:   "sorts",
		Columns: sortsColumns,
	}
}
