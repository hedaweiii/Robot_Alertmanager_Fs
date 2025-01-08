package tools

import (
	"time"
)

// ConvertUTCToChineseFormat 将 UTC 时间字符串转换为中文格式
func ConvertUTCToChineseFormat(utcTimeStr string) string {
	// 解析时间字符串
	t, _ := time.Parse(time.RFC3339Nano, utcTimeStr)

	// 转换为北京时间（UTC+8）
	beijingTime := t.In(time.FixedZone("CST", 8*3600))

	// 格式化为中文格式
	chineseFormat := beijingTime.Format("2006年01月02日 15时04分05秒")

	return chineseFormat
}
