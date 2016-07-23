package main

import (
	_ "github.com/GOCODE/pluralsight/beego/code/src/lastchance/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/public", "static")
	beego.DelStaticPath("/static")
	beego.Run()
}
