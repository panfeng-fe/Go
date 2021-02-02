package router

import (
	"iris-admin/db"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

var mssql = db.Database()

// Agent 代理模块路由
func Agent(party iris.Party) {
	var agent = party.Party("/agent")
	agent.Get("/registe", hero.Handler(Registe))
	// agent.Post("/login", hero.Handler(Login))
	// agent.Post("/logout", hero.Handler(LoginOut))
}

// Registe ...
func Registe(ctx iris.Context) {
	var res []agent
	_, err := mssql.Table(&res).Query("select top(10) AgentID,UserID,Compellation,replace(Domain,' ',''), replace(EMail,' ','') from [RYAccountsDB].[dbo].[AccountsAgent]")
	ctx.JSON(res)
	golog.Info(res, err)

}

type agent struct {
	AgentID      int
	UserID       int
	Compellation string
	Domain       string
	EMail        string
}
