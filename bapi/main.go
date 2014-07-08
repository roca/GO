package main

import (
	_ "github.com/GOCODE/bapi/docs"
	_ "github.com/GOCODE/bapi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
