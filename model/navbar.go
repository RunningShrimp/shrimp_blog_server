package model

import (
	"time"
)

type Navbar struct {
	ID          uint       `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Url         string     `json:"url" db:"url"`
	CreateTime  *time.Time `json:"create_time" db:"create_time"`
	UpdateTime  *time.Time `json:"update_time" db:"update_time"`
	DeleteTime  *time.Time `json:"delete_time" db:"delete_time"`
	Status      uint       `json:"status" db:"status"`
	NavbarLevel uint       `json:"navbar_level" db:"navbar_level"`
	ParentID    uint       `json:"parent_id" db:"parent_id"`
	Sort        uint       `json:"sort" db:"sort"`
}
