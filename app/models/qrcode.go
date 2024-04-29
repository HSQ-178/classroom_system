package models

// 二维码内容
type QrcodeInfo struct {
	Id              string `json:"id"`
	Content         string `json:"content"`
	QrcodeDuration  int    `json:"qrcodeDuration"`  //二维码有效时长
	QrcodeFrequency int    `json:"qrcodeFrequency"` //二维码刷新频率
	QrcodeExpireAt  int64  `json:"qrcodeExpireAt"`  //二维码过期时间
}
