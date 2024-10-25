package strutil

import (
	"time"
)

func StrToTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

func TimeToStr(dt time.Time) string {
	return dt.Format("2006-01-02 15:04:05")
}
