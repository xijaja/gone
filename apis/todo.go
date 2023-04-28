package apis

import (
	"github.com/gofiber/fiber/v2"
	"gone/db/model"
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
	// 将 data 颠倒
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": "2000",
		"msg":  "AllTodos",
		"data": data,
	})
}

// 更新或添加
func updateOrAddTodo(c *fiber.Ctx) error {
	// 定义请求参数结构体
	req := struct {
		Id    int    `json:"id" form:"id"`
		Title string `json:"title" form:"title" validate:"required"`
		Done  int    `json:"done" form:"done"`
	}{}
	// 绑定请求参数
	_ = c.BodyParser(&req)
	// 更新
	if req.Id != 0 {
		var todo model.Todos
		todo.FindOne(req.Id)
		todo.Title = req.Title
		todo.Done = req.Done
		todo.UpdateOne(req.Id)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": "2000",
			"msg":  "更新 todo 成功",
			"data": "",
		})
	}
	// 添加
	todo := model.Todos{
		Title: req.Title,
		Done:  req.Done,
	}
	todo.AddOne()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": "2000",
		"msg":  "添加 todo 成功",
		"data": "",
	})
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
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": "4000",
			"msg":  "id 参数有误或为空",
			"data": "",
		})
	}
	var todo model.Todos
	todo.DeleteOne(idInt)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": "2000",
		"msg":  "成功删除待办",
		"data": "",
	})
}

// 完成待办事项
func doneTodo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": "2000",
		"msg":  "ok",
		"data": "完成待办事项",
	})
}
