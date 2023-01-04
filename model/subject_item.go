package model

import (
	"time"
)

type SubjectItem struct {
	ID         uint       `json:"id" db:"id"`
	SubjectID  int        `json:"subject_id" db:"subject_id"`
	ArticleID  int        `json:"article_id" db:"article_id"`
	Status     int        `json:"status" db:"status"`
	Sort       int        `json:"sort" db:"sort"`
	CreateTime *time.Time `json:"create_time" db:"create_time"`
	UpdateTime *time.Time `json:"update_time" db:"update_time"`
	DeleteTime *time.Time `json:"delete_time" db:"delete_time"`
}
