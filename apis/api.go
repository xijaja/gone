package apis

import (
	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	api := app.Group("api") // api 路由组，访问该组路由时需加上 /api
	api.Get("/", hello)     // 保留的路由，用以验活
	api.Get("/hi", hello)   // 保留的路由，用以验活

	// 基于 /api/todos 的路由组
	todos := api.Group("todos") // api/todos 路由组
	todoApi(todos)
}

// 服务端 api 路由
func hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("👊 Yes, Iam working!")
}
