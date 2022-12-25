package model

// UserRoleDao is the manager for logic model data accessing and custom defined data operations functions management.
type UserRoleDao struct {
	Table   string          // Table is the underlying table name of the DAO.
	Group   string          // Group is the database configuration group name of current DAO.
	Columns UserRoleColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// UserRoleColumns defines and stores column names for table user_role.
type UserRoleColumns struct {
	ID         string //
	UserID     string //
	RoleID     string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// userRoleColumns holds the columns for table user_role.
var userRoleColumns = UserRoleColumns{
	ID:         "id",
	UserID:     "user_id",
	RoleID:     "role_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewUserRoleDao creates and returns a new DAO object for table data access.
func NewUserRoleDao() *UserRoleDao {
	return &UserRoleDao{
		Group:   "default",
		Table:   "user_role",
		Columns: userRoleColumns,
	}
}
