package main

import (
	"fmt"
	"log"

	"github.com/Leon1235532/test/dao"
	"github.com/Leon1235532/test/models"
	"github.com/Leon1235532/test/routers"
	"github.com/Leon1235532/test/setting"
)

const FilePath = "./config/config.ini"

func main() {
	if err := setting.Init(FilePath); err != nil {
		log.Fatalf("load mysql config failed: %#v", err.Error())
	}
	if err := dao.InitDB(setting.Conf.MySQLConfig); err != nil {
		log.Fatalf("init mysql failed: %#v", err.Error())
	}
	if err := dao.DB.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("mysql create table failed: %#v", err.Error())
	}
	defer dao.Close()
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		log.Fatalf("router register failed: %#v", err.Error())
	}
}
