package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/GOCODE/pluralsight/beego/code/src/lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["github.com/GOCODE/pluralsight/beego/code/src/lastchance/controllers:BankingController"],
		beego.ControllerComments{
			"ShowAccounts",
			`/banking`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/GOCODE/pluralsight/beego/code/src/lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["github.com/GOCODE/pluralsight/beego/code/src/lastchance/controllers:BankingController"],
		beego.ControllerComments{
			"Transfer",
			`/api/transfer`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/GOCODE/pluralsight/beego/code/src/lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["github.com/GOCODE/pluralsight/beego/code/src/lastchance/controllers:BankingController"],
		beego.ControllerComments{
			"ShowLifecycle",
			`/lifecycle`,
			[]string{"get"},
			nil})

}
