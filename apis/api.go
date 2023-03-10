package apis

import (
	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	api := app.Group("api") // api 路由组，访问该组路由时需加上 /api
	api.Get("/", hello)     // 保留的路由，用以验活
	api.Get("/hi", hello)   // 保留的路由，用以验活

	// 基于 /api/v1 的路由组
	v1 := api.Group("v1") // api/v1 路由组
	todoApi(v1)
}

// 服务端 api 路由
func hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("👊 Yes, Iam working!")
}
