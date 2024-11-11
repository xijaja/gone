package utils

import (
	"errors"
	"strings"
)

// base62Chars 包含了 Base62 编码使用的所有字符
// 顺序为 0-9, A-Z, a-z，总共 62 个字符
const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// ErrInvalidBase62 表示在解码过程中遇到了非法的 Base62 字符
var ErrInvalidBase62 = errors.New("invalid character in Base62 string")

// Base62Encode 将一个整数编码为 Base62 字符串
// Base62 编码可以将数字表示为更短的字符串，常用于短 URL 等场景
func Base62Encode(num int64) string {
	// 处理 0 的特殊情况
	if num == 0 {
		return string(base62Chars[0])
	}
	// 将负数转为正数处理，编码结果不区分正负
	if num < 0 {
		num = -num
	}
	var sb strings.Builder
	// 预分配空间，提高性能
	// 64位整数的Base62编码最长为11位 (62^11 > 2^64)
	sb.Grow(11)
	base := int64(len(base62Chars))
	// 核心编码逻辑：不断除以 62，将余数映射到 base62Chars
	for num > 0 {
		sb.WriteByte(base62Chars[num%base])
		num /= base
	}
	// 反转字符串，因为上面的过程是从低位到高位的
	bytes := []byte(sb.String())
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

// Base62Decode 将 Base62 字符串解码为整数
// 这是 Base62Encode 的逆操作
func Base62Decode(encoded string) (int64, error) {
	base := int64(len(base62Chars))
	var result int64

	// 从左到右遍历字符串，将每个字符解码并累加
	for i := 0; i < len(encoded); i++ {
		// 查找字符在 base62Chars 中的索引
		index := strings.IndexByte(base62Chars, encoded[i])
		if index < 0 {
			return 0, ErrInvalidBase62
		}
		// 累加解码结果：result = result * 62 + index
		result = result*base + int64(index)
		// 检查是否溢出
		if result < 0 {
			return 0, errors.New("Base62 decode overflow")
		}
	}
	return result, nil
}
