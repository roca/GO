package controllers

import "github.com/astaxie/beego"

// BankingController is an unexported type
type BankingController struct {
	beego.Controller
}

// URLMapping is an unexported type
func (c *BankingController) URLMapping() {
	c.Mapping("ShowAccounts", c.ShowAccounts)
}

// ShowAccounts is an unexported type
// @router /banking [get]
func (c *BankingController) ShowAccounts() {
	c.TplName = "banking.tpl"
}
