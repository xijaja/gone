package middle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

func Logs(app *fiber.App) {
	// app.Use(logger.New(logger.Config{
	// 	TimeFormat: "2006-01-02 15:04:05", // 时间格式
	// }))
	// /assets 开头的 GET 请求，不会被记录
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == http.MethodGet && len(c.Path()) >= 7 && c.Path()[0:7] == "/assets" {
			return c.Next() // 忽略，将不打印日志
		} else {
			return logger.New(logger.Config{
				TimeFormat: "2006-01-02 15:04:05", // 时间格式
			})(c) // 打印日志
		}
	})
}
