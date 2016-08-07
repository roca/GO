package controllers

import (
	"github.com/astaxie/beego"
)

// MainController  ...
type MainController struct {
	beego.Controller
}

// Get  ...
func (c *MainController) Get() {
	c.TplName = "home.tpl"
}
