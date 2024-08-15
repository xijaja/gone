package middle

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gone/db/model"
	"strconv"
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
		// 如果请求体看起来像 JSON，但 Content-Type 不正确，则设置它
		if c.Is("json") && c.Get("Content-Type") != "application/json" {
			c.Request().Header.Set("Content-Type", "application/json")
		}

		// 获取当前时间
		startTime := time.Now()
		// 在 Fiber 上下文中保存 start time
		c.Locals("startTime", startTime)

		// 保存原始请求体
		reqBody := c.Body()
		// 创建请求体的副本
		reqBodyCopy := make([]byte, len(reqBody))
		copy(reqBodyCopy, reqBody)

		// 重新设置请求体，以便后续处理器可以读取
		c.Request().SetBody(reqBodyCopy)

		// 调用下一个处理程序
		err := c.Next() // 递归调用, 从这里开始执行下一个中间件, 直到最后一个中间件, 然后再从最后一个中间件开始往前执行

		// 获取 Fiber 上下文中的 start time
		startTime, _ = c.Locals("startTime").(time.Time)
		// 获取当前时间
		endTime := time.Now()
		// 计算请求执行时间，单位为毫秒，精确到微秒
		// 例如：2.456 表示 2毫秒456微秒
		duration := float64(endTime.Sub(startTime).Microseconds()) / 1e3

		// 获取 request_id 请求 id
		rid := c.Locals("request_id").(string)

		// 解析请求体
		var reqBodyMap map[string]interface{}
		if err := json.Unmarshal(reqBody, &reqBodyMap); err != nil {
			reqBodyMap = map[string]interface{}{"error": "无法解析请求体", "original": string(reqBody)}
		}

		// 获取响应体
		resBody := c.Response().Body()
		var resBodyMap map[string]interface{}
		if err := json.Unmarshal(resBody, &resBodyMap); err != nil {
			resBodyMap = map[string]interface{}{"error": "无法解析响应体", "original": string(resBody)}
		}

		// 添加请求ID到响应体
		resBodyMap["rid"] = rid

		// 将请求存入日志库
		lg := model.Logs{
			ReqId:    rid,                                    // 请求ID
			IP:       c.Get("X-Real-IP"),                     // 请求 IP
			Url:      c.Path(),                               // 请求路径
			Method:   c.Method(),                             // 请求方法
			Status:   c.Response().StatusCode(),              // 请求状态码
			Duration: duration,                               // 请求耗时
			Params:   c.Request().URI().QueryArgs().String(), // 请求参数
			Header:   c.Request().Header.String(),            // 请求头
			Body:     reqBodyMap,                             // 请求体
			Resp:     resBodyMap,                             // 响应体
		}
		go lg.Create() // 异步入库, 提升性能且不影响主流程

		// 重写响应体
		newResBody, _ := json.Marshal(resBodyMap)
		c.Response().SetBody(newResBody)
		c.Response().Header.Set("Content-Type", "application/json")
		c.Response().Header.Set("Content-Length", strconv.Itoa(len(newResBody)))
		// 返回错误
		return err
	})
}
