package logic

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"newway/data/redis"
	"time"

	"github.com/beego/beego/v2/server/web/context"
)

type LoginInput struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	PhoneNum string `json:"phonenum"`
}

func checkprarms(ctx *context.Context) *LoginInput {
	input := &LoginInput{}
	json.Unmarshal(ctx.Input.RequestBody, input)
	return input
}

func Login(ctx *context.Context) {
	prarms := checkprarms(ctx)
	key_front := fmt.Sprintf("%s_%s", prarms.UserName, "login")
	key := getsign(key_front)
	redis.SetEx(key, prarms.UserName, 3600)

	ctx.JSONResp(key)
}

func getsign(key string) string {
	t := time.Now().Unix()
	data := fmt.Sprintf("%s_%d", key, t)
	has := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", has)
}
