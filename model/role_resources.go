package model

// RoleResourcesDao is the manager for logic model data accessing and custom defined data operations functions management.
type RoleResourcesDao struct {
	Table   string               // Table is the underlying table name of the DAO.
	Group   string               // Group is the database configuration group name of current DAO.
	Columns RoleResourcesColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// RoleResourcesColumns defines and stores column names for table role_resources.
type RoleResourcesColumns struct {
	ID         string //
	RoleID     string //
	ResourceID string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// roleResourcesColumns holds the columns for table role_resources.
var roleResourcesColumns = RoleResourcesColumns{
	ID:         "id",
	RoleID:     "role_id",
	ResourceID: "resource_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewRoleResourcesDao creates and returns a new DAO object for table data access.
func NewRoleResourcesDao() *RoleResourcesDao {
	return &RoleResourcesDao{
		Group:   "default",
		Table:   "role_resources",
		Columns: roleResourcesColumns,
	}
}
