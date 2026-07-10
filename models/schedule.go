package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `gorm:"NOT NULL" json:"title"`
	Status bool   `gorm:"DEFAULT='false'" json:"status"`
}
