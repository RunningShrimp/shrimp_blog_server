package model

import (
	"time"
)

type Comment struct {
	ID         uint       `json:"id" db:"id"`
	ArticleID  uint       `json:"article_id" db:"article_id"`
	Content    string     `json:"content" db:"content"`
	ParentID   uint       `json:"parent_id" db:"parent_id"`
	UserID     uint       `json:"user_id" db:"user_id"`
	Username   string     `json:"username" db:"username"`
	CreateTime *time.Time `json:"create_time" db:"create_time"`
	UpdateTime *time.Time `json:"update_time" db:"update_time"`
	DeleteTime *time.Time `json:"delete_time" db:"delete_time"`
	Status     uint       `json:"status" db:"status"`
}
