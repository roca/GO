package controllers

import (
	"github.com/astaxie/beego"
)

// MainController is an unexported type
type MainController struct {
	beego.Controller
}

// Get is an unexported method
func (c *MainController) Get() {
	c.TplName = "home.tpl"
}
