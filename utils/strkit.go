package utils

import (
	"encoding/json"
	"strconv"
	"time"
)

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

// 检查开始时间与结束时间是否符合逻辑
func CheckTime(startTime time.Time, endTime time.Time) bool {
	if startTime.After(endTime) && !endTime.IsZero() {
		return false
	}
	return true
}
