package controllers

import (
	"fmt"

	"github.com/GOCODE/pluralsight/beego/code/src/lastchance/models"
	"github.com/astaxie/beego"
)

// AccountController is an unexported type
type AccountController struct {
	beego.Controller
}

// Login is an unexported method
func (c *AccountController) Login() {
	if c.Ctx.Input.IsPost() {
		var loginForm models.LoginForm
		c.ParseForm(&loginForm)
		fmt.Println(loginForm)
	}
	c.TplName = "login.tpl"
}

// Create is an unexported method
func (c *AccountController) Create() {
	c.TplName = "createAccount.tpl"
}
