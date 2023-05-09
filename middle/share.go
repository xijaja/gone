package middle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CorsShare 启用跨域资源共享，这个目的是为了方便调试
func CorsShare(app *fiber.App) {
	// app.Use(cors.New()) // 默认配置
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
}
