package model

import (
	"time"
)

type Article struct {
	ID           uint       `json:"id" db:"id"`
	Title        string     `json:"title" db:"title"`
	Content      string     `json:"content" db:"content"`
	Summary      string     `json:"summary" db:"summary"`
	ClickCount   int        `json:"click_count" db:"click_count"`
	CollectCount int        `json:"collect_count" db:"collect_count"`
	FileUID      string     `json:"file_uid" db:"file_uid"`
	Recommend    uint       `json:"recommend" db:"recommend"`
	CreateTime   *time.Time `json:"create_time" db:"create_time"`
	UpdateTime   *time.Time `json:"update_time" db:"update_time"`
	DeleteTime   *time.Time `json:"delete_time" db:"delete_time"`
	Status       uint       `json:"status" db:"status"`
}
