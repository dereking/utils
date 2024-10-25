package strutil

import (
	"fmt"
	"testing"
	"time"
)

func TestRandID(t *testing.T) {
	fmt.Println(RandID())
}

func TestMD5(t *testing.T) {
	fmt.Println(MD5("123456"))
}

func TestStrToTime(t *testing.T) {
	fmt.Println(StrToTime("2018-01-01 21:59:12"))
}

func TestTimeToStr(t *testing.T) {
	fmt.Println(TimeToStr(time.Now()))
}

func TestUUID(t *testing.T) {
	fmt.Println(UUID())
}
