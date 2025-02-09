// package main

// import (
// 	"fmt"
// 	"net"
// 	"sync"
// )

// func startServer(port string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	listener, err := net.Listen("tcp", ":"+port)
// 	if err != nil {
// 		fmt.Println("Error starting server on port", port, ":", err)
// 		return
// 	}
// 	defer listener.Close()
// 	fmt.Println("Server listening on port", port)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting connection:", err)
// 			continue
// 		}
// 		go handleConnection(conn, port)
// 	}
// }

// func handleConnection(conn net.Conn, port string) {
// 	defer conn.Close()
// 	fmt.Println("New connection on port", port, "from", conn.RemoteAddr())
// }

// func main() {
// 	var wg sync.WaitGroup

// 	ports := []string{"7001", "9001"} // 두 개의 포트를 동시에 열기
// 	for _, port := range ports {
// 		wg.Add(1)
// 		go startServer(port, &wg)
// 	}

// 	wg.Wait()
// }

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTT 메시지를 저장할 변수
var lastMQTTMessage string
var mqttMutex sync.Mutex // 동시성 처리용 Mutex

// MQTT Subscriber 설정
func startMQTTSubscriber(broker, topic string) {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	client := mqtt.NewClient(opts)

	// MQTT 연결 시도
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("MQTT 연결 실패:", token.Error())
		return
	}

	fmt.Println("MQTT 구독 시작:", topic)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		mqttMutex.Lock()
		lastMQTTMessage = string(msg.Payload())
		mqttMutex.Unlock()

		// 공백 기준으로 데이터 분리
		fields := strings.Fields(lastMQTTMessage)

		// 데이터 개수가 올바른지 확인
		if len(fields) != 5 {
			fmt.Println("잘못된 메시지 형식:", lastMQTTMessage)
			return
		}

		// 각 필드 추출
		deviceID := fields[0]
		drowsiness, err1 := strconv.Atoi(fields[1])
		unintended, err2 := strconv.Atoi(fields[2])
		sudden, err3 := strconv.Atoi(fields[3])
		timestamp, err4 := strconv.ParseInt(fields[4], 10, 64)

		// 변환 오류 체크
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("데이터 변환 오류:", lastMQTTMessage)
			return
		}

		// 파싱된 데이터 출력
		fmt.Println("MQTT 메시지 수신:")
		fmt.Println("Device ID:", deviceID)
		fmt.Println("Drowsiness:", drowsiness)
		fmt.Println("Unintended:", unintended)
		fmt.Println("Sudden:", sudden)
		fmt.Println("Timestamp:", timestamp)
	})
}

// TCP 서버 실행
func startTCPServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("TCP 서버 시작 오류:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP 서버 실행 중:", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("클라이언트 연결 오류:", err)
			continue
		}
		go handleClient(conn)
	}
}

// 클라이언트 요청 처리
func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("클라이언트 연결:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("수신 데이터:", message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("데이터 읽기 오류:", err)
	}

	for {
		time.Sleep(2 * time.Second) // 주기적으로 MQTT 메시지 전송

		mqttMutex.Lock()
		msg := lastMQTTMessage
		mqttMutex.Unlock()

		if msg != "" {
			conn.Write([]byte("MQTT 메시지: " + msg + "\n"))
		}
	}
}

// 메인 함수
func main() {
	var wg sync.WaitGroup

	// MQTT Subscriber 실행 (Goroutine)
	wg.Add(1)
	go func() {
		defer wg.Done()
		startMQTTSubscriber("tcp://10.10.14.34:1883", "v2x/iot-core/accident-status")
	}()

	// TCP 서버 실행 (Goroutine)
	wg.Add(1)
	go func() {
		defer wg.Done()
		startTCPServer("7001")
	}()

	wg.Wait()
}
