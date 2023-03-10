package apis

import (
	"github.com/gofiber/fiber/v2"
)

func todoApi(v1 fiber.Router) {
	v1.Get("/todos", GetTodos)          // 获取全部 todos
	v1.Post("/todos", AddTodo)          // 添加 todo
	v1.Put("/todos/:id", UpdateTodo)    // 修改 todo
	v1.Delete("/todos/:id", DeleteTodo) // 删除 todo
}

// GetTodos 获取全部 todos
func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "AllTodos",
	})
}

// AddTodo 添加 todo
func AddTodo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "AddTodo",
	})
}

// UpdateTodo 修改 todo
func UpdateTodo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "UpdateTodo",
	})
}

// DeleteTodo 删除 todo
func DeleteTodo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "DeleteTodo",
	})
}
