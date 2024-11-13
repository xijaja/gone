package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"net/http"
	"strings"
)

func Logs(app *fiber.App) {
	// app.Use(logger.New(logger.Config{
	// 	TimeFormat: "2006-01-02 15:04:05", // 时间格式
	// }))
	// /assets 开头的 GET 请求，不会被记录
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == http.MethodGet {
			pathUrl := c.Path() // 请求路径
			// 如果请求路径中不含有 api 字样则忽略
			if !strings.Contains(pathUrl, "/api") {
				return c.Next() // 忽略，将不打印日志
			}
		}
		return logger.New(logger.Config{
			TimeFormat: "2006-01-02 15:04:05", // 时间格式
		})(c) // 打印日志
	})

	// 恐慌恢复 😱 中间件，防止程序崩溃宕机
	app.Use(recover.New())
}
