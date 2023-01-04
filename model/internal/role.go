package internal

// RoleTable is the manager for logic model data accessing and custom defined data operations functions management.
type RoleTable struct {
	Table string // Table is the underlying table name of the Table.

	Columns RoleColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// RoleColumns defines and stores column names for table role.
type RoleColumns struct {
	ID         string //
	Name       string // 角色名
	CreateTime string // 创建时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// roleColumns holds the columns for table role.
var roleColumns = RoleColumns{
	ID:         "id",
	Name:       "name",
	CreateTime: "create_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewRoleTable creates and returns a new Table object for table data access.
func NewRoleTable() *RoleTable {
	return &RoleTable{

		Table:   "role",
		Columns: roleColumns,
	}
}
