package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"restfulapi/controllers/sms"
	"restfulapi/controllers/user"
	"restfulapi/controllers/wish"
	"restfulapi/controllers/auth"
	"restfulapi/models"
)

func main() {

	//初始化DotServer
	app := dotweb.New()

	//开启debug模式
	app.SetDevelopmentMode()
	//错误处理
	app.SetExceptionHandle(func(ctx dotweb.Context, err error) {
		res := models.Res{Code: 1, Msg: "你访问的内容被互联网怪兽吃掉啦。"}
		ctx.WriteJson(res)
	})

	//设置路由
	InitRoute(app.HttpServer)

	//开始服务
	port := 80
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().POST("/SendACode", sms.SendACode)
	server.Router().POST("/Register", user.Register)
	server.Router().POST("/Login", user.Login)

	authGroup := server.Group("/Auth")
	authGroup.Use(auth.NewAccessFmtLog("user_has_token"))
	authGroup.GET("/GetMyWish", wish.MyWish)
	authGroup.POST("/AddWish", wish.AddWish)
}
