package apis

import (
	"github.com/gofiber/fiber/v2"
)

// Api 路由组，访问以下所有路由都需加上 /api
func Api(api fiber.Router) {
	api.Get("/", hello)   // 保留的路由，用以验活
	api.Get("/hi", hello) // 保留的路由，用以验活

	// 基于 /api/todos 的路由组
	todos := api.Group("todos") // api/todos 路由组
	todoApi(todos)
}

// 服务端 api 路由
func hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("👊 Yes, Iam working!")
}
