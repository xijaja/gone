package main

import (
	"github.com/gofiber/fiber/v2"
	"gone/apis"
	"gone/apis/middleware"
	"gone/config"
)

// 程序入口
func main() {
	// 创建一个 Fiber 实例
	app := fiber.New(fiber.Config{
		AppName: "Gone App", // 设置应用名称
		Prefork: *config.P,  // 是否启用多线程
	})

	// 注册路由组和中间件
	middleware.CorsShare(app) // 启用跨域资源共享，这个目的是为了方便调试
	middleware.Logs(app)      // 日志中间件，先于其他中间件，防止遗漏日志
	// 如果是生产环境
	if *config.S {
		middleware.CsrfEncrypt(app) // 启用 CSRF 保护，这个目的是为了防止跨站请求伪造
		middleware.RecordLogs(app)  // 日志入库，这个目的是为了方便查看日志
	}
	apis.Api(app.Group("api")) // 注册路由组，先于静态页面，否则其将覆盖 api
	middleware.Pages(app)      // 静态文件，将静态文件打包

	// 启动服务，监听 3030 端口
	config.GracefullyShuttingDown(app, ":3030")
}
