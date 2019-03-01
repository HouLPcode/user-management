package main

import (
	"github.com/HouLPcode/user-management/handler"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	//"github.com/stackcats/iris/middlewares"
	_ "github.com/HouLPcode/user-management/db"
)

func main() {
	//app := iris.Default()

	//runtime.GOMAXPROCS(runtime.NumCPU())
	//C.rlimit_init()

	app := iris.New()
	// 获取panic
	app.Use(recover.New())
	requestLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(requestLogger)

	//app.Use(logger.New(logger.Config{
	//	Status: true,
	//	IP:     true,
	//	Method: true,
	//	Path:   true,
	//}))

	//
	//users := app.Party("/v1/backstage/users", middlewares.ErrorHandler)
	//{
	//	users.Delete("/{user_id}", handler.DelUser)
	//}

	//系统服务
	//sd := app.Party("/sd",middlewares.ErrorHandler)
	sd := app.Party("/sd")
	{
		sd.Get("/health", handler.HandleHealth)
	}

	params := app.Party("/p")
	{
		// 127.0.0.1:50000/p/name/12
		params.Get("/{name:string}/{age}", handler.GetParamInPathStr)
		params.Get("/{id:int64}", handler.GetParamInPathInt64)
		// Querystring parameters 127.0.0.1:50000/p?p1=123
		params.Get("/", handler.GetQueryString)
		//post value
		params.Post("/form_post", handler.GotPostValue)
		// header
		params.Get("/handler_refer", handler.Refer)
	}

	cookie := app.Party("/cookies")
	{
		cookie.Get("/{name}/{value}", handler.SetMyCookie)
		cookie.Get("/{name}", handler.GETMyCookie)
		cookie.Delete("/{name}", handler.DelMyCookie)
	}

	app.Get("/panic", func(context context.Context) {
		//Internal Server Error, http 响应码 500
		panic("this is a panic")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	//app.OnErrorCode(404, func(ctx iris.Context) {
	//	ctx.JSON(iris.Map{
	//		"code":    0,
	//		"message": fmt.Sprintf("%s不存在", ctx.Path()),
	//	})
	//})

	//go func() {
	//	ch := make(chan os.Signal, 1)
	//	signal.Notify(ch,
	//		os.Interrupt,
	//		syscall.SIGINT,
	//		os.Kill,
	//		syscall.SIGKILL,
	//		syscall.SIGTERM,
	//	)
	//	select {
	//	case <-ch:
	//		println("shutdown...")
	//		timeout := 5 * time.Second
	//		ctx, cancel := context.WithTimeout(context.Background(), timeout)
	//		defer cancel()
	//		app.Shutdown(ctx)
	//	}
	//}()

	app.Run(iris.Addr(":50000"))
}
