package internal

// ArticleLabelTable is the manager for logic model data accessing and custom defined data operations functions management.
type ArticleLabelTable struct {
	Table   string              // Table is the underlying table name of the Table.
	Columns ArticleLabelColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// ArticleLabelColumns defines and stores column names for table article_label.
type ArticleLabelColumns struct {
	ID         string //
	ArticleID  string // 文章id
	LabelID    string // 标签id
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
	Status     string // 状态
}

// articleLabelColumns holds the columns for table article_label.
var articleLabelColumns = ArticleLabelColumns{
	ID:         "id",
	ArticleID:  "article_id",
	LabelID:    "label_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
	Status:     "status",
}

// NewArticleLabelTable creates and returns a new Table object for table data access.
func NewArticleLabelTable() *ArticleLabelTable {
	return &ArticleLabelTable{
		Table:   "article_label",
		Columns: articleLabelColumns,
	}
}
