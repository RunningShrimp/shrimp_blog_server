package internal

// CommentTable is the manager for logic model data accessing and custom defined data operations functions management.
type CommentTable struct {
	Table string // Table is the underlying table name of the Table.

	Columns CommentColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// CommentColumns defines and stores column names for table comment.
type CommentColumns struct {
	ID         string // 主键
	ArticleID  string // 文章id
	Content    string // 内容
	ParentID   string // 父评论id
	UserID     string // 用户id
	Username   string // 用户名
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// commentColumns holds the columns for table comment.
var commentColumns = CommentColumns{
	ID:         "id",
	ArticleID:  "article_id",
	Content:    "content",
	ParentID:   "parent_id",
	UserID:     "user_id",
	Username:   "username",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

func NewCommentDao() *CommentTable {
	return &CommentTable{
		Table:   "comment",
		Columns: commentColumns,
	}
}
