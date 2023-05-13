package middle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/utils"
	"time"
)

func CsrfEncrypt(app *fiber.App) {
	// 使用 encryptcookie 加密 cookie，用以防止 cookie 被篡改
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key:    "secret-thirty-2-character-string", // 32 字符的密钥
		Except: []string{"csrf-token"},             // 从加密中排除CSRF Cookie
	}))

	// 使用 csrf 中间件，用以防止跨站请求伪造
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "cookie:csrf-token", // 从 cookie 中获取 csrf token
		CookieName:     "csrf-token",        // csrf token 存储在 cookie 中的名称 csrf-token
		CookieSecure:   true,                // 指示 cookie 是否安全
		CookieSameSite: "Strict",            // 不允许在跨站点请求中携带 cookie
		CookieHTTPOnly: true,                // 仅允许 http 访问
		Expiration:     1 * time.Hour,       // csrf token 的过期时间
		KeyGenerator:   utils.UUID,          // 生成 csrf token 的方法
		ContextKey:     "csrf-token",        //	csrf token 存储在 ctx.Locals 中的名称
	}))
}
