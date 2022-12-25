package model

// PageTemplateDao is the manager for logic model data accessing and custom defined data operations functions management.
type PageTemplateDao struct {
	Table   string              // Table is the underlying table name of the DAO.
	Group   string              // Group is the database configuration group name of current DAO.
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

// NewPageTemplateDao creates and returns a new DAO object for table data access.
func NewPageTemplateDao() *PageTemplateDao {
	return &PageTemplateDao{
		Group:   "default",
		Table:   "page_template",
		Columns: pageTemplateColumns,
	}
}
