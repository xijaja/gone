package apis

import (
	"github.com/gofiber/fiber/v2"
	"gone/apis/handler"
	"gone/apis/middleware"
	"log"
)

// Router 路由组，访问以下所有路由都需加上 /api
func Router(app *fiber.App) {
	api := app.Group("api") // 创建 api 路由组
	api.Get("/", hello)     // 保留的路由，用以验活

	apiV1 := api.Group("/v1", middleware.Auth()) // api/v1 路由组

	var u *handler.User                // 用户管理
	apiV1.Post("/user/login", u.Login) // 登录
	apiV1.Post("/user/sth", u.PostSth) // 仅作演示

	var t *handler.Todo                   // 待办事项管理
	todos := api.Group("/todos")          // api/todos 路由组
	todos.Get("/all", t.GetAllTodos)      // api/todos/all 获取全部 todos
	todos.Post("/one", t.UpdateOrAddTodo) // api/todos/one 更新或添加
	todos.Delete("/:id", t.DeleteTodo)    // api/todos/:id 删除待办事项
	todos.Post("/done", t.DoneTodo)       // api/todos/done 完成待办事项
}

// 服务端 api 路由
func hello(c *fiber.Ctx) error {
	log.Println("hello")
	return c.Status(fiber.StatusOK).SendString("👊 Yes, Iam working!")
}
