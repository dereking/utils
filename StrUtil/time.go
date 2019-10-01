package StrUtil

import ( 
	"github.com/noaway/dateparse"
)
 

func StrToTime(str string) (time.Time ,error){
	return  dateparse.ParseAny(dateExample)
}