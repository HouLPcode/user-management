package handler

import "github.com/kataras/iris"

func HandleHealth(ctx iris.Context) {
	ctx.JSON("ok")
}
