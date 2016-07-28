package routers

import (
	"github.com/GOCODE/pluralsight/beego/code/src/lastchance/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.AccountController{})
	beego.Include(&controllers.BankingController{})
}
