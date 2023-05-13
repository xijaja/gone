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
	errs := code.Validator(req) // 验证请求参数
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(code.Bad.Reveal(fiber.Map{
			"failed": errs,
		}))
	}
	// 验证用户名和密码 fixme: 此处仅用于演示，实际应用中应该从数据库中查询并比对
	if req.Username != "admin" || req.Password != "123456" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error", "message": "用户名或密码错误", "data": nil,
		})
	}
	// todo：先将上次的 jwt 作废，再生成新的 jwt，查询数据库，如果有旧的记录则将旧的 token 作废
	// todo：注意啊：在别的需要验证 jwt 的地方，需要先确定 token 作废状态，以免自欺欺人（可以在 middle.Auth() 中间件中设置 SuccessHandler 验证通过后的逻辑）
	// 生成新的 jwt
	token := jwt.New(jwt.SigningMethodHS256)                       // 指定签名方法
	claims := token.Claims.(jwt.MapClaims)                         // 获取载荷，类型为 map[string]interface{}
	claims["username"] = req.Username                              // 存入用户名 todo: 存入用户其他信息，如权限、角色等
	claims["password"] = req.Password                              // fixme: 存入密码是不安全的，这里仅作为示例
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()          // 设置过期时间 3 天
	tokenValue, err := token.SignedString([]byte(middle.Secretly)) // 生成签名字符串
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "JWT 生成失败", "data": nil})
	}
	// todo：将 jwt 存入数据库（登录记录表），用于作废：jwt、user_id、status。
	// todo：另外，还需定时清理过期的 jwt 记录，比如 redis 拥有过期时间，可以利用 redis 的过期事件来清理
	// todo：当一个尚在有效期内的 token 被作废，则可将其存入 redis 并设置过期时间，时间可以是剩余的有效期或者覆盖，这样我们就有了一个 token 黑名单
	// todo：在 middle.Auth() 中间件中验证 token 时，先从 redis 中查询是否存在，如果存在则说明 token 已作废，然后要求用户重新登录
	// 构建返回
	return c.Status(fiber.StatusOK).JSON(code.Ok.Reveal(fiber.Map{"token": tokenValue}))
}

// 做点什么
func postSth(c *fiber.Ctx) error {
	// 从上下文 jwk 中获取 jwt
	user := c.Locals("user").(*jwt.Token) // 获取 jwt, user 是 jwt 的默认载荷名称
	// 构建返回
	return c.Status(fiber.StatusOK).JSON(
		code.Ok.Reveal(fiber.Map{
			"username": user.Claims.(jwt.MapClaims)["username"], // 获取载荷中的 username
			"password": user.Claims.(jwt.MapClaims)["password"], // 获取载荷中的 password
		}),
	)
}
