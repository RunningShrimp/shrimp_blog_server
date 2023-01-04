package internal

// LabelTable is the manager for logic model data accessing and custom defined data operations functions management.
type LabelTable struct {
	Table string // Table is the underlying table name of the Table.

	Columns LabelColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// LabelColumns defines and stores column names for table label.
type LabelColumns struct {
	ID         string //
	Name       string // 标签名
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// labelColumns holds the columns for table label.
var labelColumns = LabelColumns{
	ID:         "id",
	Name:       "name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

func NewLabelDao() *LabelTable {
	return &LabelTable{
		Table:   "label",
		Columns: labelColumns,
	}
}
