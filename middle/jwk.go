package middle

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"gone/code"
	"gone/db/dao"
	"gone/start"
	"strings"
)

// Secretly 用于生成和验证 jwt 的密钥
// var Secretly = start.Config.JwtSecret

// Auth 用于验证 jwt 的中间件
func Auth() func(ctx *fiber.Ctx) error {
	// return jwtWare.New(jwtWare.Config{
	// 	ErrorHandler: jwtError,    // 用于处理错误的函数
	// 	KeyFunc:      customKey(), // 用于解密验证的函数
	// })
	return func(ctx *fiber.Ctx) error {
		// 如果访问的是 /api/v1/user/login 则跳过验证
		if strings.HasPrefix(ctx.Path(), "/api/v1/user/login") {
			return ctx.Next()
		}
		// 返回一个中间件
		return jwtWare.New(jwtWare.Config{
			ErrorHandler:   jwtError,    // 用于处理错误的函数
			KeyFunc:        customKey(), // 用于解密验证的函数
			SuccessHandler: success,     // 验证通过后的逻辑
		})(ctx)
	}
}

// 验证通过后的逻辑
func success(c *fiber.Ctx) error {
	// 如果 token 作废，则要求用户重新登录
	token := c.Locals("user").(*jwt.Token).Raw // 获取 jwt, 并提取 token
	rds := dao.NewRedis(token)
	// 如果返回值为 1 则表示该 token 存在于黑名单之中
	if haveField := rds.IsRedisKey(); haveField == 1 {
		return c.Status(fiber.StatusUnauthorized).JSON(code.Ok.Reveal(fiber.Map{"msg": "Token 已过期，请重新登录"}))
	}
	return c.Next()
}

// 用于处理错误的函数
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(code.Bad.Reveal(fiber.Map{"msg": "缺少 Token 或格式错误"}))
	}
	return c.Status(fiber.StatusUnauthorized).JSON(code.Err.Reveal(fiber.Map{"msg": "无效或过期的 Token"}))
}

// 用于解密验证的函数
func customKey() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// 始终检查签名方法
		if t.Method.Alg() != jwtWare.HS256 {
			return nil, fmt.Errorf("以为的 jwt 签名方式 => %v", t.Header["alg"])
		}
		return []byte(start.Config.JwtSecret), nil
	}
}
