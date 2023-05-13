package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// ImReady 实例化参数
type ImReady struct {
	Url   string         // 请求地址
	Mod   string         // 请求方式
	Param map[string]any // 参数
}

// Ready 请求入口，在这里准备数据
// 返回值是 *ImReady 类型，可以链式调用
func Ready(url, mod string, can map[string]any) *ImReady {
	return &ImReady{
		Url:   url, // 请求地址
		Mod:   mod, // 请求方式
		Param: can, // 参数
	}
}

// Goo 执行请求
// 返回值：[]byte 类型，方便做后续的处理，比如 json.Unmarshal(body, &result) 绑定到结构体
// 只支持 GET/get 和 POST/post 两种请求方式，大小写都可以，其他方式请自行添加
func (ready *ImReady) Goo() []byte {
	// 传入参数
	var payload *strings.Reader
	// 判断请求方式
	switch ready.Mod {
	case "GET", "get":
		if ready.Param != nil {
			ready.Url += "?" // 当参数不为空时需要拼接参数
			for k, v := range ready.Param {
				ready.Url += fmt.Sprintf("%v=%v&", k, v)
			}
			ready.Url = ready.Url[:len(ready.Url)-1] // 拼装参数
		}
		// fmt.Println("请求地址：", r.Url)
		payload = strings.NewReader(`{}`) // GET请求时，不管参数是否为空，payload 为空即可
	case "POST", "post":
		reqBody, _ := json.Marshal(ready.Param)
		// payload := bytes.NewBuffer(reqBody)
		payload = strings.NewReader(string(reqBody)) // buffer 和 reader 都可以
	default:
		log.Println("请求模式只能是 GET 或者 POST，修改一下吧～")
		return nil
	}

	// 创建请求
	client := &http.Client{}
	req, err := http.NewRequest(ready.Mod, ready.Url, payload)
	if err != nil {
		fmt.Println("创建请求报错: ", err)
		return nil
	}
	// 添加请求头
	req.Header.Add("Content-Type", "application/json")

	// 执行请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("执行请求报错: ", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {
			fmt.Println("关闭请求报错: ", err2)
		}
	}(res.Body) // 读取完要关闭，防止泄露

	// 处理返回
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("处理返回报错: ", err)
		return nil
	}
	// 返回结果，此时 body 是 []byte 类型，可以使用 json.Unmarshal(body, &result) 转换成结构体
	return body
}
