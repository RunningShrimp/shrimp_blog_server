package model

import (
	"time"
)

type Resource struct {
	ID         uint       `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Url        string     `json:"url" db:"url"`
	CreateTime *time.Time `json:"create_time" db:"create_time"`
	UpdateTime *time.Time `json:"update_time" db:"update_time"`
	DeleteTime *time.Time `json:"delete_time" db:"delete_time"`
	Status     uint       `json:"status" db:"status"`
}
