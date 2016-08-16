package main

import (
	"github.com/GOCODE/pluralsight/beego/code/src/lastchance/filters"
	_ "github.com/GOCODE/pluralsight/beego/code/src/lastchance/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.InsertFilter("/lifecycle", beego.BeforeRouter, filters.BeforeRouterFilter)
	beego.InsertFilter("/lifecycle", beego.BeforeExec, filters.BeforeExecFilter)
	beego.InsertFilter("/lifecycle", beego.AfterExec, filters.AfterExecFilter)
	beego.InsertFilter("/lifecycle", beego.FinishRouter, filters.FinishRouterFilter)
	beego.SetStaticPath("/public", "static")
	beego.DelStaticPath("/static")
	beego.Run()
}
