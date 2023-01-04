package model

import (
	"time"
)

type Messages struct {
	ID         uint       `json:"id" db:"id"`
	Username   string     `json:"username" db:"username"`
	Content    string     `json:"content" db:"content"`
	ContactWay string     `json:"contact_way" db:"contact_way"`
	Contact    string     `json:"contact" db:"contact"`
	CreateTime *time.Time `json:"create_time" db:"create_time"`
	UpdateTime *time.Time `json:"update_time" db:"update_time"`
	DeleteTime *time.Time `json:"delete_time" db:"delete_time"`
	Status     uint       `json:"status" db:"status"`
}
