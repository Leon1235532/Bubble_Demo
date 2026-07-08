package dao

import (
	"fmt"
	"log"

	"github.com/Leon1235532/Go_backend/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB(cfg *setting.MySQLConfig) (err error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return
	} else {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Ping()
	}
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("数据库关闭失败：%#v", err)
	}
	sqlDB.Close()
}
