package auth

import (
	"rediation/controllers"
	"github.com/astaxie/beego/logs"
)

type SignController struct {
	controllers.BaseController
}

func (this *SignController) Get() {
	this.TplName = "backstage/sign-in.tpl"
}

func (this *SignController) Post() {
	form := this.Ctx.Request.Form;
	_xsrf := this.GetString("_xsrf");

	logs.GetBeeLogger().Info(_xsrf);

	this.Data["json"] = form;
	this.ServeJSON();
	this.StopRun();
}

func (this *SignController) Delete() {

}
