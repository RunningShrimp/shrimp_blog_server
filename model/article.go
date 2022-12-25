package model

// ArticleDao is the manager for logic model data accessing and custom defined data operations functions management.
type ArticleDao struct {
	Table   string         // Table is the underlying table name of the DAO.
	Group   string         // Group is the database configuration group name of current DAO.
	Columns ArticleColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// ArticleColumns defines and stores column names for table article.
type ArticleColumns struct {
	ID           string // 主键
	Title        string // 标题
	Content      string // 内容
	Summary      string // 博客简介
	ClickCount   string // 博客点击数
	CollectCount string // 博客收藏数
	FileUID      string // 标题图片uid
	Recommend    string // 是否推荐
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
	DeleteTime   string // 删除时间
	Status       string // 状态
}

// articleColumns holds the columns for table article.
var articleColumns = ArticleColumns{
	ID:           "id",
	Title:        "title",
	Content:      "content",
	Summary:      "summary",
	ClickCount:   "click_count",
	CollectCount: "collect_count",
	FileUID:      "file_uid",
	Recommend:    "recommend",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	DeleteTime:   "delete_time",
	Status:       "status",
}

// NewArticleDao creates and returns a new DAO object for table data access.
func NewArticleDao() *ArticleDao {
	return &ArticleDao{
		Group:   "default",
		Table:   "article",
		Columns: articleColumns,
	}
}
