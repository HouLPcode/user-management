package handler

import "github.com/kataras/iris"

// Set A Cookie.
func SetMyCookie(ctx iris.Context) {
	//实际使用中，cookies是存储在redis等数据库中的
	name := ctx.Params().Get("name")
	value := ctx.Params().Get("value")

	ctx.SetCookieKV(name, value)

	ctx.Writef("cookie added: %s = %s", name, value)
}

// Retrieve A Cookie.
func GETMyCookie(ctx iris.Context) {
	name := ctx.Params().Get("name")

	value := ctx.GetCookie(name)

	ctx.WriteString(value)
}

// Delete A Cookie.
func DelMyCookie(ctx iris.Context) {
	name := ctx.Params().Get("name")

	ctx.RemoveCookie(name)

	ctx.Writef("cookie %s removed", name)
}
