package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 特殊的md5 加密
func Md5Encode(mm string) string {
	h := md5.New()
	h.Write([]byte(mm))
	cipherStr := h.Sum(nil)
	tmpStr := hex.EncodeToString(cipherStr)
	return tmpStr[2:10]
}
