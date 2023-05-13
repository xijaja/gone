package middle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"net/http"
)

func Logs(app *fiber.App) {
	// app.Use(logger.New(logger.Config{
	// 	TimeFormat: "2006-01-02 15:04:05", // 时间格式
	// }))
	// /assets 开头的 GET 请求，不会被记录
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == http.MethodGet && len(c.Path()) >= 7 {
			if c.Path()[0:7] == "/assets" || c.Path()[0:11] == "/src/assets" || c.Path()[0:8] == "/favicon" {
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
