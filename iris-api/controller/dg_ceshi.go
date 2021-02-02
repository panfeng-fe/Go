package controller

import (
	"fmt"
	"github.com/kataras/iris"
)

// Testing ...
func Testing(ctx iris.Context) {
	fmt.Print(ctx.Path())
}

// PostTesting ...
func PostTesting(ctx iris.Context) {
	//fmt.Print(ctx.URLParam("agentName"))
	fmt.Print(ctx.Params().Get("agentName"))
}
