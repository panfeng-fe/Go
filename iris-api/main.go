package main

import (
	"strconv"

	"iris-api/config"
	"iris-api/controller"

	// 初始化日志
	_ "iris-api/tools/log"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		Debug:            true,
		AllowCredentials: true,
	})
	router := app.Party("/", crs).AllowMethods((iris.MethodOptions))
	{
		apiClient := router.Party("apiClient/")
		{
			apiClient.Get("video/login",controller.Testing)
			apiClient.Get("video/free",controller.Testing)
			apiClient.Get("video/detail",controller.Testing)
		}
		apiDg := router.Party("apiDg/")
		{
			apiDg.Post("account/unsettle/{agentName:string}",controller.PostTesting)
			apiDg.Post("account/order/{agentName}",controller.Testing)
			apiDg.Post("account/inform/{dg231}",controller.Testing)
			apiDg.Post("account/checkTransfer/{agentName}",controller.Testing)
			apiDg.Post("account/transfer/{agentName}",controller.Testing)
			apiDg.Post("user/getBalance/{agentName}",controller.Testing)
		}
	}
	app.Run(iris.Addr(":"+strconv.Itoa(config.Setting.App.Port)), iris.WithConfiguration(config.Setting.Project))
}

// package main

// import (
// 	"fmt"
// 	"iris-api/config"
// 	"strconv"

// 	"github.com/gohouse/converter"
// )

// func main() {
// 	var config = config.Setting.DB
// 	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.User, config.Password, config.Host, strconv.Itoa(config.Port), config.Dbname)

// 	t2t := converter.NewTable2Struct()

// 	err := t2t.
// 		SavePath("./model/struct.go").
// 		Dsn(dsn).
// 		Run()
// 	fmt.Println(err)
// }
