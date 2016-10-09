package main

import (
	"github.com/astaxie/beego"
	_ "rediation/routers"
	"github.com/astaxie/beego/logs"
)

func main() {

	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test.log","separate":["error", "warning", "notice", "info", "debug"], "level": 7}`)



	//beego.BConfig.WebConfig.TemplateLeft = "<%";
	//
	//beego.BConfig.WebConfig.TemplateRight = "%>";

	beego.Run()
}
