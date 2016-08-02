package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/GOCODE/pluralsight/beego/code/src/lastchance/models"
	"github.com/astaxie/beego"
)

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

// Transfer is an unexported type
// @router /api/transfer [post]
func (c *BankingController) Transfer() {
	var transfer models.Transfer
	json.Unmarshal(c.Ctx.Input.RequestBody, &transfer)
	fmt.Println(transfer)

	c.Ctx.WriteString("success")
}
