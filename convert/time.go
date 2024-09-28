package convert

import (
	"github.com/duke-git/lancet/v2/datetime"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai")

func Unix() int64 {
	return datetime.NewUnixNow().ToUnix()
}

func CompareBetweenInt(begin, end int64) bool {
	now := time.Now().In(loc).Unix()
	return end > begin && begin < now && end > now
}
func CompareBetweenNow(begin, end time.Time) bool {
	now := time.Now().In(loc)
	return end.After(begin) && begin.Before(now) && end.After(now)
}
func CompareNow(t time.Time) bool {
	now := time.Now().In(loc)
	return now.After(t)
}
func CompareInt(start int64) bool {
	// now 大于 起始 返回True
	now := time.Now().In(loc).Unix()
	return now > start
}
func Timestamp(timezone string) int64 {
	return datetime.Timestamp(timezone)
}
