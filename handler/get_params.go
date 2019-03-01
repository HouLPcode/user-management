package handler

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func GetParamInPathStr(ctx iris.Context) {
	var name string
	name = ctx.Params().Get("name")
	age := ctx.Params().Get("age")
	ctx.Write([]byte(name + age))
}

func GetParamInPathInt64(ctx iris.Context) {
	var id int64
	id, _ = ctx.Params().GetInt64("id") //.GetInt64Default("id",123)
	ctx.Writef("%x", id)
}

func GetQueryString(ctx iris.Context) {
	p1 := ctx.URLParam("p1")
	ctx.WriteString(p1)
}

func GotPostValue(ctx iris.Context) {
	type User struct {
		Message string `json:"message"`
		Nick    string `json:"nick"`
	}

	message := ctx.FormValue("message") //form-data  urlencode
	nick := ctx.FormValueDefault("nick", "anonymous")
	user := User{}
	ctx.ReadJSON(&user) // application/json
	//ctx.JSON(user)
	ctx.JSON(iris.Map{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

//TODO
func Refer(ctx iris.Context) {
	// request header "referer" or url parameter "referer".
	r := ctx.GetReferrer()
	switch r.Type {
	case context.ReferrerSearch:
		ctx.Writef("Search %s: %s\n", r.Label, r.Query)
		ctx.Writef("Google: %s\n", r.GoogleType)
	case context.ReferrerSocial:
		ctx.Writef("Social %s\n", r.Label)
	case context.ReferrerIndirect:
		ctx.Writef("Indirect: %s\n", r.URL)
	}
}
