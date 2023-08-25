package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gone/apis"
	"gone/middle"
	"gone/start"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 程序入口
func main() {
	// 创建一个 Fiber 实例
	app := fiber.New(fiber.Config{
		AppName: "Gone App", // 设置应用名称
		Prefork: *start.P,   // 是否启用多线程
	})

	// 注册路由组和中间件
	middle.CorsShare(app) // 启用跨域资源共享，这个目的是为了方便调试
	middle.Logs(app)      // 日志中间件，先于其他中间件，防止遗漏日志
	// 如果是生产环境
	if *start.S {
		middle.CsrfEncrypt(app) // 启用 CSRF 保护，这个目的是为了防止跨站请求伪造
		middle.RecordLogs(app)  // 日志入库，这个目的是为了方便查看日志
	}
	apis.Api(app.Group("api")) // 注册路由组，先于静态页面，否则其将覆盖 api
	middle.Pages(app)          // 静态文件，将静态文件打包

	// 启动服务，监听 3030 端口
	gracefullyShuttingDown(app, ":3030")
}

// 优雅关闭服务
func gracefullyShuttingDown(app *fiber.App, addr string) {
	// 创建一个信号量
	c := make(chan os.Signal, 1)

	// 监听中断信号
	// os.Interrupt => Ctrl+C
	// os.Kill => kill [pid]
	// syscall.SIGTERM => kill -2 [pid] => 效果类似 os.Interrupt 在程序结束之前，能够保存相关数据，然后再退出
	// syscall.SIGKILL => kill -9 [pid] => 强制终止程序，无法保存数据，直接退出，所以 signal.Notify 无法捕获该信号
	// 小贴士：为了保证程序能够正常退出，所以尽量不要使用 kill -9 [pid] 命令
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		s := <-c // 阻塞，等待中断信号
		fmt.Println("🫡 报告长官，收到信号:", s)
		_ = app.ShutdownWithTimeout(30 * time.Second) // 最多等待 30 秒
	}()

	// 启动服务，监听指定端口
	if err := app.Listen(addr); err != nil {
		log.Panic(err)
	}
	fmt.Println("🤖️ 程序已关闭...")
}
