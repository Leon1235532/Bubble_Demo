package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string `gorm:"NOT NULL" json:"title"`
	Status *bool  `gorm:"default=false" json:"status"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	// 如果Status是nil，自动设为false
	if t.Status == nil {
		falseVal := false
		t.Status = &falseVal
	}
	return nil
}
