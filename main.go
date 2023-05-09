package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gone/apis"
	"gone/middle"
	"gone/start"
)

// 程序入口
func main() {
	// 创建一个 Fiber 实例
	app := fiber.New(fiber.Config{
		AppName: "Gone App", // 设置应用名称
		Prefork: *start.P,   // 是否启用多线程
	})

	// 注册路由组和中间件
	middle.CorsShare(app)      // 启用跨域资源共享，这个目的是为了方便调试
	middle.Logs(app)           // 日志中间件，先于其他中间件，防止遗漏日志
	apis.Api(app.Group("api")) // 注册路由组，先于静态页面，否则其将覆盖 api
	middle.Pages(app)          // 静态文件，将静态文件打包

	// 启动服务，监听 3030 端口
	if err := app.Listen(":3030"); err != nil {
		fmt.Println("启动服务错误：", err)
	}
}
