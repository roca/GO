package filters

import (
	"fmt"

	"github.com/astaxie/beego/context"
)

// BeforeRouterFilter ..
func BeforeRouterFilter(ctx *context.Context) {
	fmt.Println("Before Router filter")
}

// BeforeExecFilter ...
func BeforeExecFilter(ctx *context.Context) {
	fmt.Println("Before Exec filter")
}

// AfterExecFilter ...
func AfterExecFilter(ctx *context.Context) {
	fmt.Println("After Exec filter")
}

// FinishRouterFilter ...
func FinishRouterFilter(ctx *context.Context) {
	fmt.Println("Finish Router filter")
}
