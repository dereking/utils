package strutil

import (
	"github.com/google/uuid"
)

func UUID() string {

	// 生成一个新的 UUID（UUIDv4）
	newUUID := uuid.New()
	// fmt.Println("Generated UUID:", newUUID)

	// 生成 UUID 的字符串格式
	uuidStr := newUUID.String()
	// fmt.Println("UUID as string:", uuidStr)

	return uuidStr
}

func RandID() string {
	return MD5(UUID())
}

func RandID16() string {
	return MD5(UUID())[:16]
}
