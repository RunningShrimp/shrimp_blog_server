package internal

// PageTemplateTable is the manager for logic model data accessing and custom defined data operations functions management.
type PageTemplateTable struct {
	Table string // Table is the underlying table name of the Table.

	Columns PageTemplateColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// PageTemplateColumns defines and stores column names for table page_template.
type PageTemplateColumns struct {
	ID         string //
	Name       string // 模板名
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// pageTemplateColumns holds the columns for table page_template.
var pageTemplateColumns = PageTemplateColumns{
	ID:         "id",
	Name:       "name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewPageTemplateTable creates and returns a new Table object for table data access.
func NewPageTemplateTable() *PageTemplateTable {
	return &PageTemplateTable{

		Table:   "page_template",
		Columns: pageTemplateColumns,
	}
}
