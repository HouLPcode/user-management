package handler

import (
	"github.com/HouLPcode/user-management/errno"
	"github.com/kataras/iris"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx iris.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
