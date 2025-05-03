package utils

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// 初始化日志
func InitLogger() {
	// 设置日志格式: 时间 文件 行号 日志内容
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 确保日志目录存在
	logPath := getLogFilePath()
	dir := filepath.Dir(logPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("创建日志目录失败: %v", err)
	}

	// 设置日志输出
	log.SetOutput(io.MultiWriter(
		os.Stdout, // 标准输出
		&lumberjack.Logger{
			Filename:   logPath, // 日志文件路径
			MaxSize:    10,      // 每个日志文件最大 10MB
			MaxBackups: 7,       // 保留最近 7 个备份
			MaxAge:     30,      // 保留 30 天
			Compress:   true,    // 启用压缩
		},
	))
}

// 获取日志文件路径
func getLogFilePath() string {
	// 获取年和月作为文件夹
	dir := time.Now().Format("0601")
	// 获取年月日作为文件名
	fileName := time.Now().Format("20060102")
	return "logs/" + dir + "/" + fileName + ".log"
}
