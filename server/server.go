package server

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func IrisInit() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, NotFound)
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})
	esp := app.Party("/esp", crs).AllowMethods(iris.MethodOptions)

	//传感器传入数据
	esp.Get("/sendHumiTemp", HumiTempEsp)

	app.Run(iris.Addr(":8091"))
}
