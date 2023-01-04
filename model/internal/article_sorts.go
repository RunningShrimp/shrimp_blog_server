package internal

// ArticleCategoryTable is the manager for logic model data accessing and custom defined data operations functions management.
type ArticleCategoryTable struct {
	Table   string                 // Table is the underlying table name of the Table.
	Columns ArticleCategoryColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// ArticleCategoryColumns defines and stores column names for table article_sorts.
type ArticleCategoryColumns struct {
	ID         string //
	ArticleID  string // 文章名
	SortID     string // 分类名
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// articleCategoryColumns holds the columns for table article_sorts.
var articleCategoryColumns = ArticleCategoryColumns{
	ID:         "id",
	ArticleID:  "article_id",
	SortID:     "sort_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewArticleCategoryTable creates and returns a new Table object for table data access.
func NewArticleCategoryTable() *ArticleCategoryTable {
	return &ArticleCategoryTable{
		Table:   "article_sorts",
		Columns: articleCategoryColumns,
	}
}
