package controllers

import "github.com/astaxie/beego"

// AccountController is an unexported type
type AccountController struct {
	beego.Controller
}

// Login is an unexported method
func (c *AccountController) Login() {
	c.TplName = "login.tpl"
}

// Create is an unexported method
func (c *AccountController) Create() {
	c.TplName = "createAccount.tpl"
}
