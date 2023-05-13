package middle

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gone/db/model"
	"log"
	"time"
)

// RecordLogs 将请求信息和响应信息记录到数据库
func RecordLogs(app *fiber.App) {
	// ReqId 生成 request_id 默认使用 uuid
	// 请求 ID 的选择一般与具体的业务场景有关。如果请求 ID 的主要目的是为了在日志中追踪单个请求，则生成 UUID 是更合适的选择；
	// 这是因为 UUID 具有足够的随机性和唯一性，使得每个请求 ID 都具有不同的值，这样可以防止单点故障和多次请求之间的混淆。
	// 如果请求 ID 的主要目的是为了在系统中统计和分析请求次数，或者将多个请求相关联，则使用时间戳是更合适的选择。
	// 这是因为时间戳可以提供可排序的、与时间相关的顺序，可以更好地组织请求。
	app.Use(requestid.New(requestid.Config{
		ContextKey: "request_id", // 保存 request_id 的 key
	}))

	// 计算请求处理时间
	// 追加请求 ID，重写响应体
	// 将请求信息和响应信息记录到数据库
	app.Use(func(c *fiber.Ctx) error {
		// 获取当前时间
		startTime := time.Now()
		// 在 Fiber 上下文中保存 start time
		c.Locals("startTime", startTime)
		// 调用下一个处理程序
		err := c.Next() // 递归调用, 从这里开始执行下一个中间件, 直到最后一个中间件, 然后再从最后一个中间件开始往前执行
		// 获取 Fiber 上下文中的 start time
		startTime, _ = c.Locals("startTime").(time.Time)
		// 获取当前时间
		endTime := time.Now()
		// 计算请求执行时间，秒保留 3 位小数点，时间间隔: 0.033 秒 (即：33毫秒)
		duration := float64(endTime.Sub(startTime).Milliseconds()) / 1e3

		// 获取 request_id 请求 id
		rid := c.Locals("request_id").(string)

		// 获取 API 请求体
		var reqBody map[string]interface{}
		if err := json.Unmarshal(c.Request().Body(), &reqBody); err != nil {
			reqBody = map[string]interface{}{"error": err.Error(), "original": string(c.Request().Body())}
		}
		// 获取 API 响应体
		var resBody map[string]interface{}
		if err := json.Unmarshal(c.Response().Body(), &resBody); err != nil {
			resBody = map[string]interface{}{"error": err.Error(), "original": string(c.Response().Body())}
		} else {
			resBody["rid"] = rid // 将请求 ID 存入响应体，还原响应体
			// 将响应体转为 []byte 并重写响应体
			resBodyByte, _ := json.Marshal(resBody)
			_, err = c.Write(resBodyByte)
			if err != nil {
				log.Println("重写 resBodyByte err:", err)
				return err
			}
		}

		// 请求路径
		pathUrl := c.Path()
		// 如果 pathUrl 以 /assets 或 /src/assets 或 /favicon 开头，则不记录日志
		if len(pathUrl) >= 7 {
			if pathUrl[0:7] == "/assets" || pathUrl[0:11] == "/src/assets" || pathUrl[0:8] == "/favicon" {
				return err
			}
		}

		// 将请求存入日志库
		lg := model.Logs{
			ReqId:    rid,                                    // 请求ID
			IP:       c.IP(),                                 // 请求 IP
			Url:      pathUrl,                                // 请求路径
			Method:   c.Method(),                             // 请求方法
			Status:   c.Response().StatusCode(),              // 请求状态码
			Duration: duration,                               // 请求耗时
			Params:   c.Request().URI().QueryArgs().String(), // 请求参数
			Header:   c.Request().Header.String(),            // 请求头
			Body:     reqBody,                                // 请求体
			Resp:     resBody,                                // 响应体
		}
		go lg.Create() // 异步入库, 提升性能且不影响主流程
		return err
	})
}
