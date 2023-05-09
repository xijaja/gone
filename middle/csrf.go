package middle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/utils"
	"time"
)

func CsrfEncrypt(app *fiber.App) {
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key:    "secret-thirty-2-character-string", // 32 字符的密钥
		Except: []string{"csrf-token"},             // 从加密中排除CSRF Cookie
	}))

	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-CSRF-Token", // 从表单中获取 csrf token
		CookieName:     "csrf-token",          // csrf token 存储在 cookie 中的名称 csrf-token
		CookieSameSite: "Strict",              // csrf cookie 的 SameSite 属性
		CookieHTTPOnly: true,                  // csrf cookie 是否仅允许 http 访问
		Expiration:     1 * time.Hour,         // csrf token 的过期时间
		KeyGenerator:   utils.UUID,            // 生成 csrf token 的方法
		ContextKey:     "csrf-token",          // csrf token 存储在 ctx.Locals 中的名称
	}))
}
