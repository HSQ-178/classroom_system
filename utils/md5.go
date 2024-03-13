package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(v string) string {
	d := []byte(v) //切片
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil)) // 将字节数组转换为16进制编码的字符串
}
