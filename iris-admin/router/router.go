package router

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// AllRouter 路由集合
func AllRouter(app *iris.Application) {
	// preSettring(app)
	var main = corsSetting(app)

	Agent(main)
	// Game(main)
	// Money(main)
	// Sys(main)
	// User(main)
}

func corsSetting(app *iris.Application) (main iris.Party) {
	var (
		crs context.Handler
	)

	crs = cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //允许通过的主机名称
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
		//AllowCredentials: true,
	})

	/* 定义路由 */
	main = app.Party("/", crs).AllowMethods(iris.MethodOptions)
	// main.Use(middleware.ServeHTTP)
	return main
}
