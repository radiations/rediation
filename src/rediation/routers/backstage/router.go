package backstage

import (
	"github.com/astaxie/beego"
)

func init() {
	backstageNS := beego.NewNamespace("/backstage")

	backstageNS.Namespace(authNSRoutes())

	beego.AddNamespace(backstageNS)
}
