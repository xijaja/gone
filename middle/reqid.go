package middle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// ReqId 生成 request_id 默认使用 uuid
// 请求 ID 的选择一般与具体的业务场景有关。如果请求 ID 的主要目的是为了在日志中追踪单个请求，则生成 UUID 是更合适的选择；
// 这是因为 UUID 具有足够的随机性和唯一性，使得每个请求 ID 都具有不同的值，这样可以防止单点故障和多次请求之间的混淆。
// 如果请求 ID 的主要目的是为了在系统中统计和分析请求次数，或者将多个请求相关联，则使用时间戳是更合适的选择。
// 这是因为时间戳可以提供可排序的、与时间相关的顺序，可以更好地组织请求。
func ReqId(app *fiber.App) {
	app.Use(requestid.New(requestid.Config{
		ContextKey: "request_id", // 保存 request_id 的 key
	}))
}
