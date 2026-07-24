package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string `gorm:"NOT NULL" json:"title"`
	Status *bool  `gorm:"default=false" json:"status"`
}

//GORM 的Hook，在执行 Create 保存数据到数据库之前自动调用。
func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	// 如果Status是nil，自动设为false
	if t.Status == nil {
		t.Status = new(bool)
	}
	return nil
}
