package model

import (
	"time"
)

type RoleResources struct {
	ID         uint       `json:"id" db:"id"`
	RoleID     uint       `json:"role_id" db:"role_id"`
	ResourceID uint       `json:"resource_id" db:"resource_id"`
	CreateTime *time.Time `json:"create_time" db:"create_time"`
	UpdateTime *time.Time `json:"update_time" db:"update_time"`
	DeleteTime *time.Time `json:"delete_time" db:"delete_time"`
	Status     uint       `json:"status" db:"status"`
}
