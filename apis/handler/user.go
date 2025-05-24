package handler

import (
	"errors"
	"gone/apis/middleware"
	"gone/database/cache"
	"gone/database/model"
	"gone/internal/result"
	"gone/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// User 用户管理
type User struct{}

// 登录
func (u *User) Login(c *fiber.Ctx) error {
	req := struct {
		Username string `validate:"required,min=5,max=12" json:"username"`
		Password string `validate:"required,min=6,max=18" json:"password"`
		Age      int    `validate:"required,gte=13,lte=130" json:"age"`
		Phone    string `validate:"required,len=11" json:"phone"`
	}{}
	// 绑定请求参数
	_ = c.BodyParser(&req)
	errs := utils.Validator(req) // 验证请求参数
	if errs != nil {
		return c.JSON(result.Error("请求参数错误").WithData(fiber.Map{"failed": errs}))
	}
	// 查询用户是否存在
	var user model.User
	UserResult := user.FindOneByUsername(req.Username)
	if errors.Is(UserResult.Error, gorm.ErrRecordNotFound) {
		return c.JSON(result.Error("用户不存在"))
	}
	// 验证密码是否正确
	if utils.MakeMd5(req.Password) != user.Password {
		return c.JSON(result.Error("用户或密码错误"))
	}

	// 生成新的 jwt
	claims := middleware.Claims{
		UserID:   user.Id.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)), // JWT 有效期为 7 天
			IssuedAt:  jwt.NewNumericDate(time.Now()),                         // 签发时间
		},
	}
	tokenValue, err := claims.NewJWT() // 生成签名字符串
	if err != nil {
		return c.JSON(result.Error("生成 Token 失败"))
	}
	// 构建返回
	return c.JSON(result.Success("登录成功").WithData(fiber.Map{"token": tokenValue}))
}

// Logout 登出
func (u *User) Logout(c *fiber.Ctx) error {
	// 获取 jwt 并提取 token
	token := c.Locals("user").(*jwt.Token).Raw
	// 将 token 作废，保存到 redis 标记为无效
	rds := cache.NewRedis(token)       // 将 token 作为 redis 的 key，此处 key-value 同值
	rds.SetRedisKey(token, 60*60*24*3) // 3 天后过期删除，覆盖原有的过期时间（可以通过计算设置为剩余时间，但没必要）
	return c.JSON(result.Success("登出成功"))
}

// 做点什么
func (u *User) PostSth(c *fiber.Ctx) error {
	// 从上下文 jwk 中获取 jwt
	user := c.Locals("user").(*jwt.Token) // 获取 jwt, user 是 jwt 的默认载荷名称
	// 构建返回
	return c.JSON(result.Success("登出成功").WithData(fiber.Map{
		"username":       user.Claims.(jwt.MapClaims)["username"], // 获取载荷中的 username
		"role":           user.Claims.(jwt.MapClaims)["role"],     // 获取载荷中的 role
		"json_web_token": user.Raw,                                // 获取载荷中的 json_web_token
	}))
}
