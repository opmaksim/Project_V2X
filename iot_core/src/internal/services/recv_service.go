package service

import (
	"strconv"
	"strings"
	"sync"
	"time"

	model "project/v2x/iot-core/internal/models"
	util "project/v2x/iot-core/internal/utils"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

const SERVER_PORT = 7000
const BUF_SIZE = 1024
const FIELD_CLNT = 4

// 메서드를 포함시키기 위한 구조체
type DriveEventsProcessor struct {
	recvMessage string
	recvMutex   sync.Mutex
}

type AccidentStatusProcessor struct {
	// MQTT 메시지를 저장할 변수
	// 동시성 처리용 Mutex
	recvMessage string
	recvMutex   sync.Mutex
}

//======================================================================================= DB: Drive_Events  =======================================================================================//

// func (__this *DriveEventsProcessor) createTcpServerSocket() net.Listener {
// 	// 서버 소켓 생성
// 	serverSock, err := net.Listen("tcp", ":"+strconv.Itoa(SERVER_PORT))
// 	if err != nil {
// 		log.Fatalf("Socket creation failed: %v", err)
// 		return nil
// 	}

// 	log.Infof("Server is listening on port: %v", SERVER_PORT)

// 	return serverSock
// }

// // counts of input data: 4
// func (__this *DriveEventsProcessor) parseMessage(__buff []byte, __size int) (*model.DriveEvents, error) {
// 	recvData := string(__buff[:__size])

// 	fields := strings.Fields(recvData)
// 	fmt.Printf("Length of fields: %d: ", len(fields))
// 	if len(fields) < FIELD_CLNT {
// 		log.Error("Invalid data format")
// 		return nil, errors.New("invalid data format")
// 	}

// 	deviceId := "car_01"
// 	handle, h_err := strconv.Atoi(fields[0])
// 	brake, b_err := strconv.Atoi(fields[1])
// 	accel, a_err := strconv.Atoi(fields[2])
// 	pressure, p_err := strconv.Atoi(fields[3])
// 	// Convert fields[4] to int64
// 	timeStamp := time.Now().Unix()
// 	date, d_err := util.ConvTimeToDate(timeStamp)
// 	if d_err != nil {
// 		log.Errorf("Fail to convert to date ")
// 		return nil, d_err
// 	}

// 	if h_err != nil {
// 		log.Fatal("Error: Failed to convert Handleto int")
// 		return nil, h_err
// 	}
// 	if b_err != nil {
// 		log.Fatal("Error: Failed to convert Brake to int")
// 		return nil, b_err
// 	}
// 	if a_err != nil {
// 		log.Fatal("Error: Failed to convert Accel to int")
// 		return nil, a_err
// 	}
// 	if p_err != nil {
// 		log.Fatal("Error: Failed to convert Pressure to int")
// 		return nil, a_err
// 	}

// 	dbHandler := &model.DriveEvents{
// 		DeviceId:  deviceId,
// 		Handle:    handle,
// 		Brake:     brake,
// 		Accel:     accel,
// 		Pressure:  pressure,
// 		DriveTime: date,
// 	}

// 	// device_id:

// 	return dbHandler, nil
// }

// func (__this *DriveEventsProcessor) messageArrived(__clntSock net.Conn, __db *gorm.DB) {
// 	// 소켓이 종료될 때까지 대기
// 	// C 스레드의 Join 함수와 비슷
// 	defer __clntSock.Close()

// 	//conn := &model.DBConnector{}

// 	//var dbHandler = &model.DriveEvents{}
// 	//var recvMutex = &sync.Mutex{}

// 	// 클라이언트 주소 추출
// 	clntAddr := __clntSock.RemoteAddr()
// 	log.Infof("Connection to client: %v", clntAddr)

// 	buff := make([]byte, BUF_SIZE)

// 	// 클라이언트로 부터 데이터 수신
// 	for {
// 		size, err := __clntSock.Read(buff)
// 		if err != nil {
// 			if err == io.EOF {
// 				deviceAuth.DeviceList.Disconnect(clntAddr)
// 				log.Warnf("Device %s disconnected.", clntAddr.String())
// 				break
// 			} else {
// 				log.Errorf("Recv failed from STM32: %v", err)
// 				break
// 			}
// 		}
// 		// 수신한 메시지 출력
// 		log.Infof("Received message from STRM32: %s", string(buff[:size]))

// 		// Parse
// 		dbHandler, err := __this.parseMessage(buff, size)
// 		if err != nil {
// 			log.Errorf("Fail to parse message")
// 			continue
// 		}

// 		// DB
// 		__this.recvMutex.Lock()
// 		dbHandler.InsertTables(__db)
// 		__this.recvMutex.Unlock()
// 	}
// }

// func (__this *DriveEventsProcessor) handelDeviceAuth(__clntSock net.Conn) bool {
// 	buff := make([]byte, BUF_SIZE)
// 	size, err := __clntSock.Read(buff)
// 	if err != nil {
// 		if err != io.EOF {
// 			log.Errorf("Recv failed from Device: %v", err)
// 		}
// 		return false
// 	}

// 	// 연결이 끊어졌을 경우 종료
// 	if size == 0 {
// 		log.Infof("Connection closed by device.")
// 		return false
// 	}

// 	clntAddr := __clntSock.RemoteAddr()
// 	if deviceAuth.DeviceList.Connect(__clntSock, clntAddr, string(buff[:size])) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// // go routine
// func (__this *DriveEventsProcessor) Connect(__db *gorm.DB) {
// 	serverSock := __this.createTcpServerSocket()
// 	defer serverSock.Close()

// 	// 연결
// 	for {
// 		clntSock, err := serverSock.Accept()
// 		if err != nil {
// 			log.Errorf("Handler for Drive_Events DB accept failed: %v", err)
// 			continue
// 		}

// 		if __this.handelDeviceAuth(clntSock) {
// 			// 수신 데이터 처리
// 			go __this.messageArrived(clntSock, __db)
// 		} else {
// 			clntSock.Close()
// 		}
// 	}
// }

func (__this *DriveEventsProcessor) parseMessage() *model.DriveEvents {
	// 공백 기준으로 데이터 분리
	fields := strings.Fields(__this.recvMessage)

	// 데이터 개수가 올바른지 확인
	if len(fields) != 4 {
		log.Errorf("잘못된 메시지 형식: %v", __this.recvMessage)
	}
	deviceId := "car_01"
	handle, h_err := strconv.Atoi(fields[0])
	brake, b_err := strconv.Atoi(fields[1])
	accel, a_err := strconv.Atoi(fields[2])
	pressure, p_err := strconv.Atoi(fields[3])
	// Convert fields[4] to int64
	timeStamp := time.Now().Unix()
	date, d_err := util.ConvTimeToDate(timeStamp)
	if d_err != nil {
		log.Errorf("Fail to convert to date ")
		return nil
	}

	if h_err != nil {
		log.Fatal("Error: Failed to convert Handleto int")
		return nil
	}
	if b_err != nil {
		log.Fatal("Error: Failed to convert Brake to int")
		return nil
	}
	if a_err != nil {
		log.Fatal("Error: Failed to convert Accel to int")
		return nil
	}
	if p_err != nil {
		log.Fatal("Error: Failed to convert Pressure to int")
		return nil
	}

	dbHandler := &model.DriveEvents{
		DeviceId:  deviceId,
		Handle:    handle,
		Brake:     brake,
		Accel:     accel,
		Pressure:  pressure,
		DriveTime: date,
	}
	// device_id:

	return dbHandler
}

func (__this *DriveEventsProcessor) messageArrived(__client mqtt.Client, __topic_name string, __db *gorm.DB) {
	log.Infof("Start MQTT Subscriber: %s", __topic_name)
	__client.Subscribe(__topic_name, 0, func(__client mqtt.Client, __msg mqtt.Message) {
		__this.recvMutex.Lock()
		__this.recvMessage = string(__msg.Payload())
		__this.recvMutex.Unlock()

		// Parse
		dbHandler := __this.parseMessage()

		// DB
		__this.recvMutex.Lock()
		dbHandler.InsertTables(__db)
		__this.recvMutex.Unlock()
	})
}

// MQTT Subscriber 설정
func (__this *DriveEventsProcessor) Connect(__broker_addr string, __topic_name string, __db *gorm.DB) {
	opts := mqtt.NewClientOptions().AddBroker(__broker_addr)
	client := mqtt.NewClient(opts)

	// MQTT 연결 시도
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Errorf("MQTT 연결 실패: %v", token.Error())
		return
	}

	go __this.messageArrived(client, __topic_name, __db)
}

//======================================================================================= DB: Accident_Status  =======================================================================================//

func (__this *AccidentStatusProcessor) parseMessage() *model.AccidentStatus {
	//var emotion bool
	// 공백 기준으로 데이터 분리
	fields := strings.Fields(__this.recvMessage)

	// 데이터 개수가 올바른지 확인
	if len(fields) != 1 {
		log.Errorf("잘못된 메시지 형식: %v", __this.recvMessage)
		return nil
	}

	// 각 필드 추출
	deviceID := "monitor-01"
	// drowsiness, err_d := strconv.Atoi(fields[0])
	// unintended, err_u := strconv.Atoi(fields[1])
	// sudden, err_s := strconv.Atoi(fields[2])
	emotion, err := strconv.ParseBool(fields[0])
	// 변환 오류 체크
	// if err_d != nil || err_u != nil || err_s != nil {
	// 	log.Errorf("데이터 변환 오류: %v", __this.recvMessage)
	// 	return nil
	// }
	if err != nil {
		log.Errorf("데이터 변환 오류: %v", __this.recvMessage)
	}
	// if emotionData != 0 {
	// 	emotion = true
	// } else {
	// 	emotion = false
	// }

	// 타임스탬프 발생생
	timeStamp := time.Now().Unix()
	// 변환 오류 체크
	date, t_err := util.ConvTimeToDate(timeStamp)
	if t_err != nil {
		log.Errorf("Fail to convert to date ")
	}

	// 파싱된 데이터 출력
	// fmt.Println("MQTT 메시지 수신:")
	// fmt.Println("Device ID:", deviceID)
	// fmt.Println("Drowsiness:", drowsiness)
	// fmt.Println("Unintended:", unintended)
	// fmt.Println("Sudden:", sudden)
	// fmt.Println("Timestamp:", timestamp)

	// dbHandler := &model.AccidentStatus{
	// 	DeviceId:   deviceID,
	// 	Drowsiness: drowsiness,
	// 	Unintended: unintended,
	// 	Sudden:     sudden,
	// 	EventTime:  date,
	// }
	dbHandler := &model.AccidentStatus{
		DeviceId:  deviceID,
		Emotion:   emotion,
		EventTime: date,
	}

	return dbHandler
}

func (__this *AccidentStatusProcessor) messageArrived(__client mqtt.Client, __topic_name string, __db *gorm.DB) {
	log.Infof("Start MQTT Subscriber: %s", __topic_name)
	__client.Subscribe(__topic_name, 0, func(__client mqtt.Client, __msg mqtt.Message) {
		__this.recvMutex.Lock()
		__this.recvMessage = string(__msg.Payload())
		__this.recvMutex.Unlock()

		// Parse
		dbHandler := __this.parseMessage()

		// DB
		__this.recvMutex.Lock()
		dbHandler.InsertTables(__db)
		__this.recvMutex.Unlock()
	})
}

// MQTT Subscriber 설정
func (__this *AccidentStatusProcessor) Connect(__broker_addr string, __topic_name string, __db *gorm.DB) {
	opts := mqtt.NewClientOptions().AddBroker(__broker_addr)
	client := mqtt.NewClient(opts)

	// MQTT 연결 시도
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Errorf("MQTT 연결 실패: %v", token.Error())
		return
	}

	go __this.messageArrived(client, __topic_name, __db)
}
