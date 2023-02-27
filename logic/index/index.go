package index

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web/context"
)

type IndexInput struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	PhoneNum string `json:"phonenum"`
}

func icheckprarms(ctx *context.Context) *IndexInput {
	input := &IndexInput{}
	json.Unmarshal(ctx.Input.RequestBody, input)
	return input
}

func Index(ctx *context.Context) {
	ctx.JSONResp("index")
}
