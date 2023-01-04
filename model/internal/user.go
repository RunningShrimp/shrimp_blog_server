package internal

// UserTable is the manager for logic model data accessing and custom defined data operations functions management.
type UserTable struct {
	Table string // Table is the underlying table name of the Table.

	Columns UserColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	ID         string // 主键
	Username   string // 用户名
	Password   string // 密码
	Email      string // 邮箱
	Avatar     string // 头像
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// userColumns holds the columns for table user.
var userColumns = UserColumns{
	ID:         "id",
	Username:   "username",
	Password:   "password",
	Email:      "email",
	Avatar:     "avatar",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewUserTable creates and returns a new Table object for table data access.
func NewUserTable() *UserTable {
	return &UserTable{
		Table:   "user",
		Columns: userColumns,
	}
}