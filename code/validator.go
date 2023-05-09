package code

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gone/utils"
	"log"
)

// ErrorResponse 错误响应
type ErrorResponse struct {
	Field    string `json:"field"`     // 失败字段
	Tag      string `json:"tag"`       // 标签
	Value    string `json:"value"`     // 值
	ErrorMsg string `json:"error_msg"` // 错误信息
}

// Validator 参数验证器
func Validator(st interface{}) []*ErrorResponse {
	var validate = validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(st)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = utils.CamelToSnake(err.StructNamespace())
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.ErrorMsg = validatorErrorMsgMaker(element.Field, element.Tag, element.Value)
			errors = append(errors, &element)
		}
	}
	return errors
}

// validatorErrorMsgMaker 验证错误信息生成器
func validatorErrorMsgMaker(field, tag, value string) string {
	switch tag {
	case "len":
		return fmt.Sprintf("字段 %s，长度必须为 %s 位", field, value)
	case "min":
		return fmt.Sprintf("字段 %s，长度至少为 %s 位", field, value)
	case "max":
		return fmt.Sprintf("字段 %s，长度至多为 %s 位", field, value)
	case "eq":
		return fmt.Sprintf("字段 %s，必须等于 %s", field, value)
	case "ne":
		return fmt.Sprintf("字段 %s，不能等于 %s", field, value)
	case "gt":
		return fmt.Sprintf("字段 %s，必须大于 %s", field, value)
	case "lt":
		return fmt.Sprintf("字段 %s，必须小于 %s", field, value)
	case "gte":
		return fmt.Sprintf("字段 %s，必须大于或等于 %s", field, value)
	case "lte":
		return fmt.Sprintf("字段 %s，必须小于或等于 %s", field, value)
	case "contains":
		return fmt.Sprintf("字段 %s，必须包含 %s", field, value)
	case "excludes":
		return fmt.Sprintf("字段 %s，不能包含 %s", field, value)
	case "required":
		return fmt.Sprintf("字段 %s，不能为空", field)
	case "oneof":
		return fmt.Sprintf("字段 %s，必须为 %s 中的一个", field, value)
	case "email":
		return fmt.Sprintf("字段 %s，必须为邮箱格式", field)
	case "iscolor":
		return fmt.Sprintf("字段 %s，必须为颜色 hexcolor|rgb|rgba|hsl|hsla 格式", field)
	case "rgb":
		return fmt.Sprintf("字段 %s，必须为 RGB 格式", field)
	case "rgba":
		return fmt.Sprintf("字段 %s，必须为 RGBA 格式", field)
	case "ip":
		return fmt.Sprintf("字段 %s，必须为 IP 地址格式", field)
	case "url":
		return fmt.Sprintf("字段 %s，必须为 URL 格式", field)
	case "file":
		return fmt.Sprintf("字段 %s，必须是文件", field)
	case "dir":
		return fmt.Sprintf("字段 %s，必须是文件夹", field)
	case "number":
		return fmt.Sprintf("字段 %s，必须为数字格式", field)
	case "json":
		return fmt.Sprintf("字段 %s，必须为 JSON 格式", field)
	case "uuid":
		return fmt.Sprintf("字段 %s，必须为 UUID 格式", field)
	case "md5":
		return fmt.Sprintf("字段 %s，必须为 MD5 格式", field)
	case "sha256":
		return fmt.Sprintf("字段 %s，必须为 SHA256 格式", field)
	case "base64":
		return fmt.Sprintf("字段 %s，必须为 Base64 格式", field)
	case "alpha":
		return fmt.Sprintf("字段 %s，必须为字母格式", field)
	case "alphanum":
		return fmt.Sprintf("字段 %s，必须为字母或数字格式", field)
	case "startswith":
		return fmt.Sprintf("字段 %s，必须以 %s 开头", field, value)
	case "endswith":
		return fmt.Sprintf("字段 %s，必须以 %s 结尾", field, value)
	case "startnotwith":
		return fmt.Sprintf("字段 %s，不能以 %s 开头", field, value)
	case "endnotwith":
		return fmt.Sprintf("字段 %s，不能以 %s 结尾", field, value)
	default:
		log.Println(fmt.Sprintf("未知的字段验证错误: 字段: %s 标签: %s 值: %s", field, tag, value))
		return fmt.Sprintf("字段 %s，验证失败", field)
	}
}
