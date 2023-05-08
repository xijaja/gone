package code

import (
	"github.com/gofiber/fiber/v2"
)

// 自定义返回格式
type code uint

// 全局 code 码
const (
	Ok  = code(2000) // 请求成功
	Bad = code(4000) // 请求错误
	Err = code(5000) // 系统异常
)

// Reveal 流露：按特定格式返回数据并将请求信息存入数据库
func (co code) Reveal(c *fiber.Ctx, resp fiber.Map) (turn fiber.Map) {
	rid := c.Locals("request_id").(string) // 获取 request_id
	msg := "success"
	// 如果有 msg 信息则将 resp 中的 msg 移到上层，否则返回默认值
	if resp["msg"] != nil {
		msg = resp["msg"].(string)
		delete(resp, "msg")
	} else if co >= Bad {
		msg = "error"
	}
	// 构建返回
	return fiber.Map{
		"code": co,   // 返回状态码
		"msg":  msg,  // 返回消息
		"rid":  rid,  // 请求ID
		"data": resp, // 返回数据
	}
}
