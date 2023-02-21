package middleware

import (
	"fmt"
	"newway/data/redis"

	"github.com/beego/beego/v2/server/web/context"
)

type ErrorResp struct {
	Erron  int         `json:"erron"`
	Errmsg string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func CheckAuth(ctx *context.Context) {
	sign := ctx.GetCookie("sign")
	if len(sign) == 0 {
		ctx.JSONResp("need a sign")
	}
	_, err := redis.Get(sign)

	if err != nil {
		resp := &ErrorResp{
			Erron:  101,
			Errmsg: "auth error",
		}
		ctx.JSONResp(resp)
	}

	fmt.Println("before")

}
