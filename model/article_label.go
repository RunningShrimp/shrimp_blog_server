package model

import (
	"time"
)

type ArticleLabel struct {
	ID         uint       `json:"id" db:"id"`
	ArticleID  uint       `json:"article_id" db:"article_id"`
	LabelID    uint       `json:"label_id" db:"label_id"`
	CreateTime *time.Time `json:"create_time" db:"create_time"`
	UpdateTime *time.Time `json:"update_time" db:"update_time"`
	DeleteTime *time.Time `json:"delete_time" db:"delete_time"`
	Status     uint       `json:"status" db:"status"`
}
