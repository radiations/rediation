package models

import (
	//_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"time"
	"github.com/satori/go.uuid"
)

var engine *xorm.Engine

/**
 * 获取当前 xorm引擎实例（根据配置初始化的）
 */
func GetEngine() *xorm.Engine {
	return engine;
}

/**
 * 注册所有模型
 * @param engine xorm引擎实例
 */
func registerModels(engine *xorm.Engine) {
	engine.Sync2(new(User))
}


func init() {

	driverName := beego.AppConfig.DefaultString("db::driverName", "mysql")
	datasource := beego.AppConfig.DefaultString("db::datasource", "root:root@tcp(127.0.0.1:3306)/rediation?charset=utf8&loc=Asia%2FShanghai")
	maxIdleConn := beego.AppConfig.DefaultInt("db::maxIdleConn", 10)
	maxOpenConn := beego.AppConfig.DefaultInt("db::maxOpenConn", 30)

	var err error

	engine, err = xorm.NewEngine(driverName, datasource)

	if nil != err {
		logs.GetBeeLogger().Error("Fail to create xorm system logger: %v\n", err)
	}

	engine.SetMaxIdleConns(maxIdleConn)
	engine.SetMaxOpenConns(maxOpenConn)

	engine.ShowSQL(true)
	engine.ShowExecTime(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)

	registerModels(engine)

}

type BaseModel struct {
	Id			string		`xorm:"pk varchar(36) notnull"`
	CreateTime	int64		//`xorm:"create"`
	CreateIp	string		`xorm:"varchar(63) notnull"`
	UpdateTime	int64		//`xorm:"updated"`
	UpdateIp	string		`xorm:"varchar(63)"`
}

func (model *BaseModel) BeforeInsert () {
	model.Id = uuid.NewV1().String()
	model.CreateTime = time.Now().UnixNano()
}

func (model BaseModel) BeforeUpdate () {
	model.UpdateTime = time.Now().UnixNano()
}
