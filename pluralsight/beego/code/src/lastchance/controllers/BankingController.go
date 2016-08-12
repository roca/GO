package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/GOCODE/pluralsight/beego/code/src/lastchance/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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

	response := models.TransferResponse{
		Transfer: transfer,
	}

	if isValid {
		response.Status = "success"
	} else {
		response.Status = "failure"
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// ShowLifecycle ...
// @router /lifecycle [get]
func (c *BankingController) ShowLifecycle() {
	fmt.Println("Action Execution")
}

// Init ...
func (c *BankingController) Init(ctx *context.Context, controllerName string, actionName string, app interface{}) {
	fmt.Printf("Initilaization: %s.%s\n", controllerName, actionName)
	c.Controller.Init(ctx, controllerName, actionName, app)
}

// Prepare ...
func (c *BankingController) Prepare() {
	fmt.Println("Prepare controller")
	c.Controller.Prepare()
}

// Render ...
func (c *BankingController) Render() error {
	fmt.Println("Render result")
	c.Ctx.WriteString("result")
	return nil
}

// Finish ...
func (c *BankingController) Finish() {
	fmt.Println("Finish controller")
	c.Controller.Finish()
}
