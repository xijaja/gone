package utils

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"time"
)

// 初始化日志
func InitLogger() {
	// 设置日志格式: 时间 文件 行号 日志内容
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 设置日志输出
	log.SetOutput(io.MultiWriter(
		os.Stdout, // 标准输出
		&lumberjack.Logger{
			Filename:   getLogFilePath(), // 日志文件路径
			MaxSize:    10,               // 每个日志文件最大 10MB
			MaxBackups: 7,                // 保留最近 7 个备份
			MaxAge:     30,               // 保留 30 天
			Compress:   true,             // 启用压缩
		},
	))
}

// 获取日志文件路径
func getLogFilePath() string {
	today := time.Now().Format("2006-01-02")
	return "logs/" + today + ".log"
}
