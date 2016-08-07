package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/GOCODE/pluralsight/beego/code/src/lastchance/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// BankingController ...
type BankingController struct {
	beego.Controller
}

// URLMapping ...
func (c *BankingController) URLMapping() {
	c.Mapping("ShowAccounts", c.ShowAccounts)
}

// ShowAccounts ...
// @router /banking [get]
func (c *BankingController) ShowAccounts() {
	c.Data["accounts"] = []models.Account{
		models.Account{
			ID:     1,
			Name:   "Checking",
			Number: "8888",
			Amount: 642.27,
		},
		models.Account{
			ID:     2,
			Name:   "Savings",
			Number: "3344",
			Amount: 1000,
		},
	}

	c.TplName = "banking.tpl"
}

// Transfer ...
// @router /api/transfer [post]
func (c *BankingController) Transfer() {
	var transfer models.Transfer
	json.Unmarshal(c.Ctx.Input.RequestBody, &transfer)
	fmt.Println(transfer)

	valid := validation.Validation{}
	isValid, _ := valid.Valid(&transfer)
	fmt.Println(transfer)
	fmt.Println(valid.ErrorMap())

	var message string
	if isValid {
		message = "success"
	} else {
		message = "failure"
	}

	c.Ctx.WriteString(message)
}
