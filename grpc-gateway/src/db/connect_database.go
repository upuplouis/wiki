package db

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var DB *gorm.DB

func init()  {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		logrus.Fatal(err)
	}
	sqlUser := cfg.Section("mysql").Key("db.User").String()
	sqlPasswd := cfg.Section("mysql").Key("db.Pwd").String()
	sqlIp := cfg.Section("mysql").Key("db.Host").String()
	sqlPort := cfg.Section("mysql").Key("db.Port").String()
	sqlDatabase := cfg.Section("mysql").Key("db.Name").String()

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlUser, sqlPasswd, sqlIp, sqlPort, sqlDatabase)
	DB, err = gorm.Open("mysql", conn)
	logrus.Info("db ok")
	if err != nil {
		logrus.Fatal("connect database error", err)
	}
	DB.DB().SetMaxIdleConns(16)
	DB.DB().SetMaxOpenConns(128)
}
