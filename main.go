package main

import (
	redis "newway/data/redis"
	_ "newway/router"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func Initlog() {
	logs.SetLevel(logs.LevelDebug)
	logs.SetLogger(logs.AdapterFile, `{"filename":"./log/app.log", "level":6}`)
	logs.EnableFuncCallDepth(true)
}

func InitAll() {
	redis.InitRedis()
	Initlog()
}

func main() {
	InitAll()
	//web.AutoRouter(&UserController{})
	web.Run()
}
