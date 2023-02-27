package router

import (
	"newway/logic/index"
	"newway/logic/login"
	"newway/middleware"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	v1 := web.NewNamespace("/v1",
		web.NSBefore(middleware.CheckAuth),
		web.NSGet("index", index.Index),
	)
	v2 := web.NewNamespace("v2", web.NSPost("login", login.Login))

	web.AddNamespace(v1)
	web.AddNamespace(v2)
	//v1.Filter("befor", web.MiddleWare)

}

func Index(ctx *context.Context) {
	ctx.JSONResp("hello")
}
