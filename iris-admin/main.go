package main

import (
	"iris-admin/conf"
	"iris-admin/router"
	"strconv"

	"github.com/kataras/iris"
	_"iris-admin/tools"
)

func main() {
	app := iris.New()
	router.AllRouter(app)

	app.Run(iris.Addr(":"+strconv.Itoa(conf.Setting.App.Port)), iris.WithConfiguration(conf.Setting.Project))

}

// app.RegisterView(iris.HTML("resources", ".html").Binary(static.Asset, static.AssetNames))
// staticHandler := iris.StaticEmbeddedHandler("resources", static.Asset, static.AssetNames, false)
// app.SPA(staticHandler).AddIndexName("index.html")
