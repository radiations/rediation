package main

import (
//	"github.com/astaxie/beego"
	_ "rediation/routers"
	"github.com/astaxie/beego/logs"
	. "rediation/models"
	//"github.com/satori/go.uuid"
)


func main() {

	//logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"], "level": 7}`)
	logs.SetLogger(logs.AdapterConsole, `{"level":1}`)

	engine := GetEngine()

	sess := engine.NewSession()

	defer sess.Close()

	sess.Begin()
	user := new(User)
	//user.Id = uuid.NewV1().String()
	user.Name = "test"

	//id, err := sess.InsertMulti([]*models.User{user})
	id, err := sess.Insert(user)


	if nil != err {
		logs.GetBeeLogger().Error("插入数据失败", err)
		return
	}

	logs.GetBeeLogger().Info("插入数据成功：%s", id)


	sess.Commit()
	//
	//
	//fmt.Println(o.Insert(user))

	//beego.Run()
}
