package utils

import "github.com/google/uuid"

// IsUUID 判断 string 是否为 UUID 格式
// 如果是，则返回 uuid.UUID 类型的 UUID 和 true
func IsUUID(str string) (u uuid.UUID, isu bool) {
	// 判断长度
	if len(str) != 32 {
		return uuid.Nil, false
	}
	// 是否能转为 UUID
	u, err := uuid.Parse(str)
	if err != nil {
		return uuid.Nil, false
	}
	return u, true
}

// StringToUUID 将 string 转为 uuid.UUID 类型
func StringToUUID(str string) (uuid.UUID, error) {
	return uuid.Parse(str)
}
