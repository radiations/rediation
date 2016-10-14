package backstage

import (
	"github.com/astaxie/beego"
	"rediation/controllers/backstage/auth"
)



func authNSRoutes() *beego.Namespace {
	authNS := beego.NewNamespace("/auth")

	authNS.Router("/sign", &auth.SignController{})

	return authNS
}

