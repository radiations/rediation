package routers

import (
	"rediation/controllers"
	"github.com/astaxie/beego"
)

func init() {



    beego.Router("/", &controllers.MainController{})
}
