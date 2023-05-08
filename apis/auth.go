package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gone/code"
	"gone/middle"
	"time"
)

// 登录
func login(c *fiber.Ctx) error {
	req := struct {
		Username string `validate:"required,min=5,max=12" json:"username"`
		Password string `validate:"required,min=6,max=18" json:"password"`
		Age      int    `validate:"required,gte=13,lte=130" json:"age"`
		Phone    string `validate:"required,len=11" json:"phone"`
	}{}
	// 绑定请求参数
	_ = c.BodyParser(&req)
	errs := middle.ParameterValidator(req) // 验证请求参数
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(code.Bad.Reveal(c, fiber.Map{
			"failed": errs,
		}))
	}
	// 验证用户名和密码 fixme: 此处仅用于演示，实际应用中应该从数据库中查询并比对
	if req.Username != "admin" || req.Password != "123456" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "用户名或密码错误", "data": nil,
		})
	}
	// 生成 jwt
	token := jwt.New(jwt.SigningMethodHS256)                       // 指定签名方法
	claims := token.Claims.(jwt.MapClaims)                         // 获取载荷，类型为 map[string]interface{}
	claims["username"] = req.Username                              // 存入用户名 todo: 存入用户其他信息，如权限、角色等
	claims["password"] = req.Password                              // fixme: 存入密码是不安全的，这里仅作为示例
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()          // 设置过期时间
	tokenValue, err := token.SignedString([]byte(middle.Secretly)) // 生成签名字符串
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "JWT 生成失败", "data": nil})
	}
	// 构建返回
	return c.Status(fiber.StatusOK).JSON(code.Ok.Reveal(c, fiber.Map{"token": tokenValue}))
}

// 做点什么
func postSth(c *fiber.Ctx) error {
	// 从上下文 jwk 中获取 jwt
	user := c.Locals("user").(*jwt.Token) // 获取 jwt, user 是 jwt 的默认载荷名称
	return c.Status(fiber.StatusOK).JSON(
		code.Ok.Reveal(c, fiber.Map{
			"username": user.Claims.(jwt.MapClaims)["username"], // 获取载荷中的 username
			"password": user.Claims.(jwt.MapClaims)["password"], // 获取载荷中的 password
		}),
	)
}
