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
	r, err := redis.Get(sign)
	if err != nil {
		fmt.Println(r)
		panic(err)
	}

	if r != nil {
	} else {
		resp := &ErrorResp{
			Erron:  101,
			Errmsg: "auth error",
		}
		ctx.JSONResp(resp)
	}

	fmt.Println("before")

}
