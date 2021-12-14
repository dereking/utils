package StrUtil

import (
	// "github.com/noaway/dateparse"
	"time"
)

func StrToTime(str string) (time.Time, error) {
	//return dateparse.ParseAny(str), nil
	return time.Now(), nil
}
