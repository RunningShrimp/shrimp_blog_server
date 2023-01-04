package internal

// MessagesTable is the manager for logic model data accessing and custom defined data operations functions management.
type MessagesTable struct {
	Table string // Table is the underlying table name of the Table.

	Columns MessagesColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// MessagesColumns defines and stores column names for table messages.
type MessagesColumns struct {
	ID         string //
	Username   string // 昵称
	Content    string // 内容
	ContactWay string // 联系途径
	Contact    string // 联系方式
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// messagesColumns holds the columns for table messages.
var messagesColumns = MessagesColumns{
	ID:         "id",
	Username:   "username",
	Content:    "content",
	ContactWay: "contact_way",
	Contact:    "contact",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewMessagesTable creates and returns a new Table object for table data access.
func NewMessagesTable() *MessagesTable {
	return &MessagesTable{

		Table:   "messages",
		Columns: messagesColumns,
	}
}
