package utils

import (
	"bytes"
	"encoding/json"
	"log"
)

// PrettyStruct 格式化结构体
func PrettyStruct(data interface{}) string {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println("PrettyStruct 格式化结构体错误:", err)
	}
	return string(val)
}

// PrettyString 格式化字符串
func PrettyString(str string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		log.Println("PrettyString 格式化字符串错误:", err)
		return ""
	}
	return prettyJSON.String()
}

// Struct2JsonString 将结构体实例转换为JSON字符串
func Struct2JsonString(data interface{}) string {
	val, err := json.Marshal(data)
	if err != nil {
		log.Println("Struct2JsonString 将结构体实例转换为JSON字符串:", err)
	}
	return string(val)
}
