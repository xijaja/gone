package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gone/apis"
	"gone/db/access"
	"gone/db/model"
	"gone/middle"
	"gone/start"
)

// 初始化数据库
func init() {
	// 自动迁移，入参如 &model.Movie{}, &model.Todos{}
	err := access.DB.AutoMigrate(&model.Todos{})
	if err != nil {
		fmt.Println("数据库迁移失败:", err)
	} else {
		fmt.Println("数据库初始化完成")
	}
}

// 程序入口
func main() {
	// 创建一个 Fiber 实例
	app := fiber.New(fiber.Config{
		AppName: "Gone App", // 设置应用名称
		Prefork: *start.P,   // 是否启用多线程
	})

	// 启用跨域资源共享，这个目的是为了方便调试
	// app.Use(cors.New()) // 默认配置
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// 注册路由组和中间件
	middle.Logs(app)  // 日志中间件，先于其他中间件，防止遗漏日志
	apis.Api(app)     // 注册路由组，先于静态页面，否则其将覆盖 api
	middle.Pages(app) // 静态文件，将静态文件打包

	// 启动服务，监听 3030 端口
	if err := app.Listen(":3030"); err != nil {
		fmt.Println("启动服务错误：", err)
	}
}
