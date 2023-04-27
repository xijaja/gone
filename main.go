package main

import (
	"github.com/gofiber/fiber/v2"
	"gone/apis"
	"gone/middle"
	"gone/start"
)

func main() {
	// 创建一个 Fiber 实例
	app := fiber.New(fiber.Config{
		AppName: "Gone App", // 设置应用名称
		Prefork: *start.P,   // 是否启用多线程
	})

	// 注册路由组
	// 必须先于 middle.Pages 注册否则其将覆盖 api
	apis.Api(app)

	// 注册中间件
	middle.Logs(app)  // 日志中间件
	middle.Pages(app) // 静态文件中间件

	// 启动服务，监听 3000 端口
	_ = app.Listen(":3000")
}
