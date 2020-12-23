package util

import "time"

// EndTimeOfDay 获取t时间当晚零点
func EndTimeOfDay(t time.Time) time.Time {
	beginTimeOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return beginTimeOfDay.Add(time.Hour * 24)
}
