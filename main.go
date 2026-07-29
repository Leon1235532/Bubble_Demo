package main

import (
	"fmt"
	"log"

	"github.com/Leon1235532/GoTask/auth"
	"github.com/Leon1235532/GoTask/dao"
	"github.com/Leon1235532/GoTask/models"
	"github.com/Leon1235532/GoTask/routers"
	"github.com/Leon1235532/GoTask/setting"
)

const FilePath = "./config/config.ini"

func main() {
	if err := setting.Init(FilePath); err != nil {
		log.Fatalf("load mysql config failed: %#v", err.Error())
	}
	auth.InitJwt(setting.Conf.JwtSecret)

	if err := dao.InitDB(setting.Conf.MySQLConfig); err != nil {
		log.Fatalf("init mysql failed: %#v", err.Error())
	}

	if err := dao.DB.AutoMigrate(&models.Todo{}, &models.User{}); err != nil {
		log.Fatalf("create some table failed: %#v", err.Error())
	}
	r := routers.SetupRouter()

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		log.Fatalf("router register failed: %#v", err.Error())
	}
	defer dao.Close()
}
