package utils

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

var gcfg *ini.File

func init()  {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		logrus.Fatal(err)
	}
	gcfg = cfg
}

func GetValueFromIni(section string, key string) string {
	return gcfg.Section(section).Key(key).String()
}