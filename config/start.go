package config

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gone/internal/auto"
	"gone/pkg/utils"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

// ---------------------------------------------
// 启动配置
// ---------------------------------------------

// S P T 启动变量
var S = flag.Bool("s", false, "true 为生产环境，默认 false 开发环境")
var P = flag.Bool("p", false, "true 为启用多线程，默认 false 不启动")
var T = flag.Bool("t", false, "true 为启动定时任务，默认 false 不启动")

// Config 初始化配置
var Config = projectConfig{}

// 初始化配置信息
func init() {
	// 初始化日志
	utils.InitLogger()

	// 检查是否在测试环境中
	if testing.Testing() {
		log.Println("当前为🧪测试环境🧪")
		// 在测试环境中使用测试专用配置
		Config = Config.getMyConfig(false) // 使用开发环境配置
		return
	}

	// 解析命令行参数
	flag.Parse()

	// 设置为发布模式
	if *S {
		Config = Config.getMyConfig(true) // 赋值为生产环境配置
		log.Printf("当前为🔥生产环境🔥 定时任务启动状态:%v\n", *T)
	} else {
		Config = Config.getMyConfig(false) // 赋值为开发环境配置
		log.Printf("当前为🌲开发环境🌲 定时任务启动状态:%v\n", *T)
	}

	// 启动定时任务
	if *T {
		go auto.ScheduledTasks()
	}
}

// ---------------------------------------------
// 关闭服务
// ---------------------------------------------

// GracefullyShuttingDown 优雅关闭服务
func GracefullyShuttingDown(app *fiber.App, addr string) {
	// 创建一个信号量
	c := make(chan os.Signal, 1)

	// 监听中断信号
	// os.Interrupt => Ctrl+C
	// os.Kill => kill [pid]
	// syscall.SIGTERM => kill -2 [pid] => 效果类似 os.Interrupt 在程序结束之前，能够保存相关数据，然后再退出
	// syscall.SIGKILL => kill -9 [pid] => 强制终止程序，无法保存数据，直接退出，所以 signal.Notify 无法捕获该信号
	// 小贴士：为了保证程序能够正常退出，所以尽量不要使用 kill -9 [pid] 命令
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// 创建一个同步等待组
	var serverShutdown sync.WaitGroup
	go func() {
		s := <-c
		fmt.Println("🫡 报告长官，收到信号:", s)
		serverShutdown.Add(1)                         // 通知 WaitGroup，当前 goroutine 已启动
		defer serverShutdown.Done()                   // 通知 WaitGroup，当前 goroutine 已结束
		_ = app.ShutdownWithTimeout(10 * time.Second) // 最多等待 10 秒
	}() // 阻塞，等待中断信号

	// 获取监听地址
	if addr == "" {
		addr = ":3030" // 如果 addr 为空，则默认为 :3030
	} else if addr[0] != ':' {
		addr = ":" + addr // 如果 addr 开头不是 :，则添加 :
	}

	// 启动服务，监听指定端口
	if err := app.Listen(addr); err != nil {
		log.Panic(err)
	}
	// 等待所有服务关闭
	serverShutdown.Wait()
	fmt.Println("🤖️ 程序已关闭...")
}
