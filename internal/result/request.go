package result

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ImReady 实例化参数
type ImReady struct {
	Url     string            // 请求地址
	Method  string            // 请求方法
	Body    map[string]any    // 参数
	Headers map[string]string // 自定义请求头
}

// Ready 请求入口，在这里准备数据
func Ready(method, url string) *ImReady {
	return &ImReady{
		Url:     url,
		Method:  method,
		Headers: map[string]string{},
	}
}

// SetBody 设置请求参数
func (r *ImReady) SetBody(body map[string]any) *ImReady {
	r.Headers["Content-Type"] = "application/json"
	r.Body = body
	return r
}

// SetForm 设置表单参数
func (r *ImReady) SetForm(form map[string]string) *ImReady {
	r.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	for k, v := range form {
		r.Body[k] = v
	}
	return r
}

// SetHeader 设置请求头
func (r *ImReady) SetHeader(key, value string) *ImReady {
	r.Headers[key] = value
	return r
}

// Goo 执行请求
func (r *ImReady) Goo() ([]byte, error) {
	var payload io.Reader

	// 处理请求参数
	switch r.Method {
	case http.MethodGet:
		if r.Body != nil {
			params := url.Values{}
			for k, v := range r.Body {
				params.Add(k, fmt.Sprint(v))
			}
			r.Url += "?" + params.Encode()
		}
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		if r.Body != nil {
			jsonBody, err := json.Marshal(r.Body)
			if err != nil {
				return nil, fmt.Errorf("参数序列化错误: %w", err)
			}
			payload = bytes.NewBuffer(jsonBody)
		}
	}

	// 创建请求
	req, err := http.NewRequest(r.Method, r.Url, payload)
	if err != nil {
		return nil, fmt.Errorf("创建请求错误: %w", err)
	}

	// 设置请求头
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// 执行请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("执行请求错误: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	// 读取响应
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应错误: %w", err)
	}

	// 返回响应
	return body, nil
}

// 使用示例
// body, err := Ready(http.MethodPost, "https://api.example.com/data").
//     SetBody(map[string]any{"key": "value"}).
//     SetHeader("Authorization", "Bearer token").
//     Goo()
// if err != nil {
//     log.Printf("请求错误: %v", err)
//     return
// }

// 处理响应
// var result SomeStruct
// if err := json.Unmarshal(body, &result); err != nil {
//     log.Printf("解析响应错误: %v", err)
//     return
// }
