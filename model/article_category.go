package model

import (
	"time"
)

type ArticleCategory struct {
	ID         uint       `json:"id" db:"id"`
	ArticleID  int        `json:"article_id" db:"article_id"`
	CategoryID int        `json:"category_id" db:"category_id"`
	CreateTime *time.Time `json:"create_time" db:"create_time"`
	UpdateTime *time.Time `json:"update_time" db:"update_time"`
	DeleteTime *time.Time `json:"delete_time" db:"delete_time"`
	Status     uint       `json:"status" db:"status"`
}
