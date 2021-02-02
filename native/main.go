package main

import (
	// 初始化日志
	_ "native/tools/log"
	"strconv"

	"native/config"
	"native/controller"
	_ "native/model"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
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
		index := router.Party("index/")
		{
			index.Get("index", controller.Index)
			index.Post("login", controller.Login)
		}

		user := router.Party("user/")
		{
			user.Get("player", controller.Info)
			user.Get("player/{id :int}", controller.Find)

			user.Get("online", controller.Info)
			user.Get("online/{id :int}", controller.Find)

			user.Get("game", controller.Info)
			user.Get("game/{id :int}", controller.Find)

			user.Post("player/modify", controller.PlayerModify)
			user.Get("player/profit", controller.Testing)
			user.Get("player/bet", controller.Testing)
			user.Post("player/forbid", controller.PlayerForbid)
			user.Post("player/delBackCard", controller.Testing)
			user.Post("player/userInfo", controller.PlayerUserInfo)
		}

		recharge := router.Party("recharge/")
		{
			recharge.Get("list", controller.Info)
			recharge.Get("list/{id :int}", controller.Find)

			recharge.Get("recharge", controller.Info)
			recharge.Get("recharge/{id :int}", controller.Find)

			recharge.Get("withdrawal", controller.Info)
			recharge.Get("withdrawal/{id :int}", controller.Find)

			recharge.Get("present", controller.Info)
			recharge.Get("present/{id :int}", controller.Find)

			recharge.Get("list/msg", controller.ListMsg)
		}

		agent := router.Party("agent/")
		{
			agent.Get("managent", controller.Testing)
			agent.Get("managent/{id :int}", controller.Testing)

			agent.Post("add/addAgent", controller.Testing)
			agent.Post("add/addUser", controller.Testing)

			agent.Post("managent/forbid", controller.Testing)
			agent.Post("managent/modify", controller.Testing)
		}

		chat := router.Party("chat/")
		{
			chat.Get("info", controller.Info)
			chat.Get("info/{id :int}", controller.Find)

			chat.Post("unMsg", controller.Testing)
			chat.Post("goChat", controller.Testing)
			chat.Post("submit", controller.Testing)
		}

		pay := router.Party("pay/")
		{
			pay.Get("callback", controller.Testing)
			pay.Post("agree", controller.Testing)
			pay.Post("reject", controller.Testing)
			pay.Post("artificial", controller.Testing)
		}

		sys := router.Party("sys/")
		{
			sys.Get("admin", controller.Info)
			sys.Get("admin/{id :int}", controller.Find)

			sys.Post("admin/add", controller.Testing)
			sys.Post("admin/del", controller.Testing)
			sys.Post("admin/submit", controller.Testing)

			sys.Get("station", controller.Testing)
			sys.Post("station/set", controller.Testing)

			sys.Get("affiche", controller.Testing)
			sys.Post("affiche/set", controller.Testing)
		}
	}
	app.Run(iris.Addr(":"+strconv.Itoa(config.Setting.App.Port)), iris.WithConfiguration(config.Setting.Project))
}

// package main

// import (
// 	"fmt"
// 	"native/config"
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
