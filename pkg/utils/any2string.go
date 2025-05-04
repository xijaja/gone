package utils

import (
	"encoding/json"
	"fmt"
	"github.com/nilorg/sdk/convert"
	"log"
)

// AnyToString 类型转换工具，去吧：字符串
func AnyToString(src any) string {
	if src == nil {
		fmt.Println("src为空")
	}
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return convert.ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		log.Println("AnyToString 类型转换失败", err)
		return ""
	}
	return string(data)
}
