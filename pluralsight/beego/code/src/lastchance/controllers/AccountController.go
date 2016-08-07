package controllers

import (
	"fmt"

	"github.com/GOCODE/pluralsight/beego/code/src/lastchance/models"
	"github.com/astaxie/beego"
)

// AccountController  ...
type AccountController struct {
	beego.Controller
}

// Login  ...
func (c *AccountController) Login() {
	if c.Ctx.Input.IsPost() {
		var loginForm models.LoginForm
		c.ParseForm(&loginForm)
		fmt.Println(loginForm)
	}
	c.TplName = "login.tpl"
}

// Create  ...
func (c *AccountController) Create() {
	c.TplName = "createAccount.tpl"
}
