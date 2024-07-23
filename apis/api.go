package apis

import (
	"github.com/gofiber/fiber/v2"
	"gone/middle"
)

// Api 路由组，访问以下所有路由都需加上 /api
func Api(api fiber.Router) {
	api.Get("/", hello)   // 保留的路由，用以验活
	api.Get("/hi", hello) // 保留的路由，用以验活

	api.Post("/login", login)                // 登录 fixme: 仅作演示
	api.Post("/sth", middle.Auth(), postSth) // 带有权限验证 fixme: 仅作演示

	// 基于 /api/todos 的路由组
	todos := api.Group("todos")         // api/todos 路由组
	todos.Get("/all", getAllTodos)      // 获取全部 todos
	todos.Post("/one", updateOrAddTodo) // 更新或添加
	todos.Delete("/:id", deleteTodo)    // 删除待办事项
	todos.Post("/done", doneTodo)       // 完成待办事项
}

// 服务端 api 路由
func hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("👊 Yes, Iam working!")
}
