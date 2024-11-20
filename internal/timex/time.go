package timex

import (
	"time"
)

// TimeNowFormat 格式化当前时间
func TimeNowFormat() string {
	return Format(time.Now(), "2006-01-02 15:04:05")
}

// TimeNowUnix 获取当前时间戳
func TimeNowUnix() int64 {
	return time.Now().Unix()
}

// Format 将时间转换为字符串
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}
