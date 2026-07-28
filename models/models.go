package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:20;unique;not null" json:"username"`
	Pwd      string `gorm:"size:60;not null" json:"-"`
	Todos    []Todo `gorm:"foreignKey:UID"`
}

type Todo struct {
	gorm.Model
	Title  string `gorm:"not null" json:"title"`
	Status *bool  `json:"status"`
	UID    uint   `gorm:"index" json:"uid"`
	User   *User  `gorm:"foreignKey:UID" json:"-"`
}

//GORM 的Hook，在执行 Create 保存数据到数据库之前自动调用。
func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	// 如果Status是nil，自动设为false
	if t.Status == nil {
		t.Status = new(bool)
	}
	return nil
}
