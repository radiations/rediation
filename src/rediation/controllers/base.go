package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"strings"
	"html/template"
)



type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {

	//CSRF防范的表单
	this.Data["CSRFInput"] = template.HTML(this.XSRFFormHTML())
	//设置当前路径相对项目部署的根路径
	this.Data["RelativeToBase"] = relativeBase(this.Ctx.Input.URL()[1:])

}

var relativeReg *regexp.Regexp = regexp.MustCompile(`[^/]+`)

func relativeBase(url string) string {
	url = url[strings.Index(url, "/") + 1:]

	//logs.GetBeeLogger().Info(url)
	return relativeReg.ReplaceAllLiteralString(url, "..")
}