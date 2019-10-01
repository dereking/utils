package StrUtil

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/noaway/dateparse"
)

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}


func StrToTime(str string) (time.Time ,error){
	return  dateparse.ParseAny(dateExample)
}