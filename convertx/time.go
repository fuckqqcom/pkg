package convertx

import (
	"github.com/duke-git/lancet/v2/datetime"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai")

func Unix() int64 {
	return datetime.NewUnixNow().ToUnix()
}

func CompareInt(begin, end int64) bool {
	now := time.Now().In(loc).Unix()
	return end > begin && begin < now && end > now
}
func CompareNow(begin, end time.Time) bool {
	now := time.Now().In(loc)
	return end.After(begin) && begin.Before(now) && end.After(now)
}

func Timestamp(timezone string) int64 {
	return datetime.Timestamp(timezone)
}
