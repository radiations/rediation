package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/config"
)

type Config struct {
	Key 		string		`xorm:"varchar(50) pk notnull"`
	Value 		string		`xorm:"varchar(255)"`
}

const (
	CONFIG_PRIVATE_KEY 		= 	"PRIVATE_KEY"
	CONFIG_PUBLIC_KEY		=	"PUBLIC_KEY"
)

func GetConfig(key string, value string) string {

	config := Config{Key:key}

	has, err := engine.Get(&config)

	if !has {
		logs.GetBeeLogger().Error("获取配置信息出错%s", err)
		return nil
	}

	logs.GetBeeLogger().Info("获取配置信息成功%s", config)

	return config.Value
 }

func SetConfig(key string, value string) string {



	return nil;
}