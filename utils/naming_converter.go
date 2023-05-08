package utils

import "unicode"

// CamelToSnake 驼峰到蛇形命名转换器
func CamelToSnake(str string) string {
	runes := []rune(str)
	var newRunes []rune
	for idx, r := range runes {
		if unicode.IsUpper(r) && idx > 0 {
			newRunes = append(newRunes, '_')
		}
		newRunes = append(newRunes, unicode.ToLower(r))
	}
	return string(newRunes)
}
