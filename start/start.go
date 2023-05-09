package start

import (
	"flag"
	"fmt"
	"gone/auto"
	"gone/utils"
)

// ---------------------------------------------
// 启动配置
// ---------------------------------------------

// S P T 启动变量
var S = flag.Bool("s", false, "true 为生产环境，默认 false 开发环境")
var P = flag.Bool("p", false, "true 为启用多线程，默认 false 不启动")
var T = flag.Bool("t", false, "true 为启动定时任务，默认 false 不启动")
var B = flag.Bool("b", false, "true 为执行前端构建，默认 false 不构建")

// Config 初始化配置
var Config = projectConfig{}

// 初始化配置信息
func init() {
	// 解析命令行参数
	flag.Parse()

	// 设置为发布模式
	if *S {
		Config = Config.getMyConfig(true) // 赋值为生产环境配置
		fmt.Printf("当前为🔥生产环境🔥 定时任务启动状态:%v\n", *T)
	} else {
		Config = Config.getMyConfig(false) // 赋值为开发环境配置
		fmt.Printf("当前为🌲开发环境🌲 定时任务启动状态:%v\n", *T)
	}

	// 执行编译前端的命令
	if *B {
		utils.RunCmd(fmt.Sprintf("cd ./frontend && %s run build", Config.PkgManager))
	}

	// 启动定时任务
	if *T {
		go auto.ScheduledTasks()
	}
}
