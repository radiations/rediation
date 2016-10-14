package routers

import (
	"rediation/controllers"
	"github.com/astaxie/beego"
	_ "rediation/routers/backstage"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
