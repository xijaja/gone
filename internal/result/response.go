package result

import (
	"github.com/gofiber/fiber/v2"
)

// Res 响应结构体
type Res struct {
	Code uint      `json:"code"`
	Msg  string    `json:"msg"`
	Data fiber.Map `json:"data"`
}

// 请求成功
func Success(msg ...string) *Res {
	if len(msg) > 0 {
		return &Res{2000, msg[0], fiber.Map{}}
	}
	return &Res{2000, "success", fiber.Map{}}
}

// 请求错误
func Error(msg ...string) *Res {
	if len(msg) > 0 {
		return &Res{4000, msg[0], fiber.Map{}}
	}
	return &Res{4000, "error", fiber.Map{}}
}

// 无权限
func NoPermission(msg ...string) *Res {
	if len(msg) > 0 {
		return &Res{4001, msg[0], fiber.Map{}}
	}
	return &Res{4001, "no permission", fiber.Map{}}
}

// 系统异常
func SystemError(msg ...string) *Res {
	if len(msg) > 0 {
		return &Res{5000, msg[0], fiber.Map{}}
	}
	return &Res{5000, "system error", fiber.Map{}}
}

// WithData 设置响应数据
func (r Res) WithData(data fiber.Map) *Res {
	return &Res{
		Code: r.Code,
		Msg:  r.Msg,
		Data: data,
	}
}
