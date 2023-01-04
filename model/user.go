package model

import (
	"time"
)

type User struct {
	ID         uint       `json:"id" db:"id"`
	Username   string     `json:"username" db:"username"`
	Password   string     `json:"password" db:"password"`
	Email      string     `json:"email" db:"email"`
	Avatar     string     `json:"avatar" db:"avatar"`
	CreateTime *time.Time `json:"create_time" db:"create_time"`
	UpdateTime *time.Time `json:"update_time" db:"update_time"`
	DeleteTime *time.Time `json:"delete_time" db:"delete_time"`
	Status     uint       `json:"status" db:"status"`
}
