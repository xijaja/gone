package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"gone/config"
	"gone/database/dao"
	"gone/internal/code"
	"strings"
	"time"
)

// NewJWT 用于创建 jwt 的中间件
func NewJWT(username, role string, days uint8) (tokenValue string, err error) {
	tokenSign := jwt.New(jwt.SigningMethodHS256)
	claims := tokenSign.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(days)).Unix()
	return tokenSign.SignedString([]byte(config.Config.JwtSecret)) // JwtSecret 生成 jwt 的密钥
}

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
	token := c.Locals("user").(*jwt.Token) // 获取 jwt, 并提取 token
	rds := dao.NewRedis(token.Raw)
	// 如果返回值为 1 则表示该 token 存在于黑名单之中
	if haveField := rds.IsRedisKey(); haveField == 1 {
		return c.Status(fiber.StatusUnauthorized).JSON(code.Red.Reveal(fiber.Map{"msg": "Token 已过期，请重新登录"}))
	}
	// 自动续期
	// 如果 token 的有效期小于 7 天，则修改 token 的有效期加 1 天
	nowTime := time.Now().Unix()
	if token.Valid && token.Claims.(jwt.MapClaims)["exp"].(float64)-float64(nowTime) < float64(60*60*24*7) {
		token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(time.Hour * 24).Unix()
	} else {
		return c.JSON(code.Red.Reveal(fiber.Map{"msg": "登录过期"}))
	}
	return c.Next()
}

// 用于处理错误的函数
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(code.Red.Reveal(fiber.Map{"msg": "缺少 Token 或格式错误"}))
	}
	return c.Status(fiber.StatusUnauthorized).JSON(code.Red.Reveal(fiber.Map{"msg": "无效或过期的 Token"}))
}

// 用于解密验证的函数
func customKey() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// 始终检查签名方法
		if t.Method.Alg() != jwtWare.HS256 {
			return nil, fmt.Errorf("以为的 jwt 签名方式 => %v", t.Header["alg"])
		}
		return []byte(config.Config.JwtSecret), nil // JwtSecret 验证 jwt 的密钥
	}
}
