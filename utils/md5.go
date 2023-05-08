package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MakeMd5 获取MD5加密字符串
func MakeMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
