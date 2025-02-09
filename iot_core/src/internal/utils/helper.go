package utils

import (
	_ "project/v2x/iot-core/internal/auths"
	"time"
)

func ConvTimeToDate(__timestamp int64) (string, error) {
	date := time.Unix(__timestamp, 0)

	// Format the time to "YYYY-MM-DD HH:MM:SS"
	// This is no hard coding
	return string(date.Format("2006-01-02 15:04:05")), nil
}

// 문자열을 고정 크기 바이트 배열로 변환하는 함수
// func copyId(target [auth.ID_SIZE]byte, src string) [auth.ID_SIZE]byte {
// 	copy(target[:], src)
// 	return target
// }
