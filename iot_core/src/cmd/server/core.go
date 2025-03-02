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
	"os"
	"os/signal"
	controller "project/v2x/iot-core/internal/controllers"
	"syscall"

	conn "project/v2x/iot-core/internal/models"

	log "github.com/sirupsen/logrus"
)

// 메인 함수
func main() {
	mariaDB := conn.InitDBInfo()
	db, err := mariaDB.Connect()
	if err != nil {
		log.Fatal("Fail to connect to DataBase.")
	}

	mariaDB.CreateTables(db)

	accHandle := &controller.AccidentStatusHandler{}
	driveHandle := &controller.DriveEventsHandler{}
	sender := &controller.SendHandler{}

	accHandle.Handler(db)
	driveHandle.Handler(db)
	sender.Handler(db)

	// for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// wait shutdown when occure signal via sigChan
	select {
	case sig := <-sigChan:
		log.Infof("Receive signal: %s, Shutting down...\n", sig)
	}

	// 모든 핸들러가 완료될 때까지 대기
	//controller.WaitForHandlers()
	mariaDB.Disconnect()
}
