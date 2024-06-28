package apis

import (
	"github.com/gofiber/fiber/v2"
	"gone/code"
	"gone/db/model"
	"gone/middle"
)

func todoApi(todos fiber.Router) {
	todos.Get("/all", getAllTodos)      // 获取全部 todos
	todos.Post("/one", updateOrAddTodo) // 更新或添加
	todos.Delete("/:id", deleteTodo)    // 删除待办事项
	todos.Post("/done", doneTodo)       // 完成待办事项
}

// 获取全部 todos
func getAllTodos(c *fiber.Ctx) error {
	var todos model.Todos
	var data = todos.FindAll()
	return c.Status(fiber.StatusOK).JSON(code.Ok.Reveal(c, fiber.Map{
		"list": data,
	}))
}

// 更新或添加
func updateOrAddTodo(c *fiber.Ctx) error {
	// 定义请求参数结构体
	req := struct {
		Id    int    `json:"id" form:"id"`
		Title string `json:"title" form:"title"`
		Done  bool   `json:"done" form:"done"`
	}{}
	// 绑定请求参数
	_ = c.BodyParser(&req)
	// 验证请求参数
	if errs := middle.ParameterValidator(req); errs != nil {
		return c.Status(fiber.StatusOK).JSON(code.Bad.Reveal(c, fiber.Map{"failed": errs}))
	}
	// 更新
	if req.Id != 0 {
		var todo model.Todos
		todo.FindOne(req.Id)
		todo.Title = req.Title
		todo.Done = &req.Done
		todo.UpdateOne(req.Id)
		return c.Status(fiber.StatusOK).JSON(code.Ok.Reveal(c, fiber.Map{
			"msg": "更新 todo 成功",
		}))
	}
	// 添加
	todo := model.Todos{
		Title: req.Title,
		Done:  &req.Done,
	}
	todo.AddOne()
	return c.Status(fiber.StatusOK).JSON(code.Ok.Reveal(c, fiber.Map{
		"msg": "添加 todo 成功",
	}))
}

// 删除待办事项
func deleteTodo(c *fiber.Ctx) error {
	// 获取路由参数
	idStr := c.Params("id")
	// 将字符串转换为 int
	var idInt int
	for _, v := range idStr {
		idInt = idInt*10 + int(v-'0')
	}
	if idInt == 0 {
		return c.Status(fiber.StatusOK).JSON(code.Bad.Reveal(c, fiber.Map{
			"msg": "id 参数有误或为空",
		}))
	}
	var todo model.Todos
	todo.DeleteOne(idInt)

	return c.Status(fiber.StatusOK).JSON(code.Ok.Reveal(c, fiber.Map{
		"msg": "成功删除待办",
	}))
}

// 完成待办事项
func doneTodo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(code.Ok.Reveal(c, fiber.Map{
		"msg": "该接口尚未完善",
	}))
}
