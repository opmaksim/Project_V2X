package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	serverAddr := "10.10.14.54:7000"

	// 무한 루프를 돌며 서버에 지속적으로 데이터를 전송
	for {
		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			fmt.Println("서버 연결 실패:", err)
			time.Sleep(3 * time.Second) // 연결 재시도 대기
			continue
		}

		fmt.Println("서버 연결 성공:", serverAddr)

		// Auth device
		_, err = conn.Write([]byte("[PJW_TEST:PASSWD]"))

		for {
			// 4개의 랜덤 숫자 생성
			data := []string{}
			for i := 0; i < 4; i++ {
				data = append(data, strconv.Itoa(rand.Intn(1000))) // 0~999 범위 숫자
			}

			// 공백으로 구분하여 문자열 생성
			message := fmt.Sprintf("%s\n", joinWithSpace(data))

			// 서버로 데이터 전송
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println("데이터 전송 실패:", err)
				break
			}

			fmt.Println("전송된 데이터:", message)

			time.Sleep(2 * time.Second) // 2초 대기 후 다음 데이터 전송
		}

		conn.Close()
	}
}

// 슬라이스를 공백으로 구분하여 문자열로 변환하는 함수
func joinWithSpace(data []string) string {
	result := ""
	for i, v := range data {
		if i > 0 {
			result += " "
		}
		result += v
	}
	return result
}
