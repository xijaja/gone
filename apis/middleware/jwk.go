package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gone/internal/config"
	"log"
	"time"
)

// Claims 定义 JWT 的声明
type Claims struct {
	UserID   string `json:"user_id"`  // 用户ID
	Username string `json:"username"` // 用户邮箱
	jwt.RegisteredClaims
}

// NewJWT 用于创建 jwt 的中间件
func (claims *Claims) NewJWT() (tokenString string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(config.Config.JwtSecret))
}

// JwtAuth 用于验证 jwt 的中间件
func JwtAuth(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		log.Println("缺少 Token 或格式错误")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or malformed token"})
	}

	// Token 通常以 "Bearer <token>" 的形式出现
	const BearerSchema = "Bearer "
	if len(authHeader) <= len(BearerSchema) || authHeader[:len(BearerSchema)] != BearerSchema {
		log.Println("token 格式错误")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}
	tokenString := authHeader[len(BearerSchema):]

	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// 确保 token 的签名算法是我们期望的
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("token 签名算法错误")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.Config.JwtSecret, nil
	})
	if err != nil {
		log.Println("token 解析错误")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token: " + err.Error()})
	}

	// 检查 token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 检查 JWT 是否过期
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() >= int64(exp) {
				log.Println("token 已过期")
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token is expired"})
			}
		} else {
			log.Println("token 过期时间缺失")
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims: exp missing"})
		}
		ctx.Locals("user", token) // 将解析后的 token 存储在 c.Locals 中，后续处理函数可以使用
		return ctx.Next()
	}

	// 如果 token 无效，则返回 401 错误
	log.Println("token 无效, 需要重新登录")
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
}
