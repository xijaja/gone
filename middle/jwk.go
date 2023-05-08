package middle

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// Secretly 用于生成和验证 jwt 的密钥
var Secretly = "qU2CMS2ZV0Vb4RU30ILavs@kmNCvm8x4d2t2J"

// Auth 用于验证 jwt 的中间件
func Auth() func(ctx *fiber.Ctx) error {
	return jwtWare.New(jwtWare.Config{
		ErrorHandler: jwtError,    // 用于处理错误的函数
		KeyFunc:      customKey(), // 用于解密验证的函数
	})
}

// 用于处理错误的函数
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "缺少 JWT 或格式错误", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "无效或过期的 JWT", "data": nil})
}

// 用于解密验证的函数
func customKey() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// 始终检查签名方法
		if t.Method.Alg() != jwtWare.HS256 {
			return nil, fmt.Errorf("以为的 jwt 签名方式 => %v", t.Header["alg"])
		}
		return []byte(Secretly), nil
	}
}
