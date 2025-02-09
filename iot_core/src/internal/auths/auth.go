package auth

import (
	_ "errors"
	"net"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

const DEVICE_MAX_CLNT = 32
const ID_SIZE = 20

type DeviceAccount struct {
	Index int      // Index는 int로 정의
	Conn  net.Conn // fd도 int로 정의
	Ip    net.Addr // IP는 고정 크기 배열로 정의
	Id    string   // id는 고정 크기 배열로 정의
	Pw    string   // pw도 고정 크기 배열로 정의
}

type DeviceAuth struct {
	authMutex sync.Mutex
	Device    [DEVICE_MAX_CLNT]DeviceAccount
}

var DeviceList DeviceAuth

// 디바이스 인증 함수
func (__this *DeviceAuth) Connect(__deviceConn net.Conn, __deviceAddr net.Addr, __deviceInfo string) bool {
	var deviceCnt uint8 = 0
	//__this.authMutex = &sync.Mutex{}

	// ':' 또는 '[' 또는 ']'로 문자열을 분리
	tokens := strings.FieldsFunc(__deviceInfo, func(r rune) bool {
		return r == ':' || r == '[' || r == ']'
	})

	if len(tokens) < 2 {
		log.Errorf("Invalid accounts format for author device")
		return false
	}

	deviceID := tokens[0]
	devicePW := tokens[1]

	log.Infof("Tokenized data: deviceID = %s, devicePW = %s", deviceID, devicePW)

	// 인증 실패 메시지는 최초 한 번만 출력
	//authFailed := false

	// 클라이언트 연결 상태를 확인
	for i := 0; i < DEVICE_MAX_CLNT; i++ {
		if DeviceList.Device[i].Id == deviceID {
			// 이미 접속된 경우
			if DeviceList.Device[i].Conn != nil {
				log.Warnf("[%s] Already logged", deviceID)
				return false
			}

			// 비밀번호 확인
			if DeviceList.Device[i].Pw == devicePW {
				// 인증 성공, 정보 갱신
				__this.authMutex.Lock()
				DeviceList.Device[i].Index = i
				DeviceList.Device[i].Conn = __deviceConn
				DeviceList.Device[i].Ip = __deviceAddr
				deviceCnt++
				__this.authMutex.Unlock()

				log.Infof("[%s] Authentication successful (ip:%s, count of devices:%d)", deviceID, __deviceAddr.String(), deviceCnt)

				return true
			}
		}
	}

	// 인증 실패 시 메시지 출력 (한 번만 출력)
	//if !authFailed {
	log.Errorf("Authentication failed! Device not found: [%s]", deviceID)
	//authFailed = true
	//}

	return false
}

// 연결이 끊어진 디바이스 정리
func (__this *DeviceAuth) Disconnect(__deviceAddr net.Addr) {
	var deviceCnt uint8 = 0
	//__this.authMutex = &sync.Mutex{}

	for i := 0; i < DEVICE_MAX_CLNT; i++ {
		if DeviceList.Device[i].Ip == __deviceAddr {
			log.Warnf("Disconnecting device: %s\n", DeviceList.Device[i].Id)

			// 파일 디스크립터 및 연결 정보 초기화
			__this.authMutex.Lock()
			DeviceList.Device[i].Conn = nil
			DeviceList.Device[i].Ip = nil
			DeviceList.Device[i].Index = 0

			// 전체 연결된 디바이스 수 감소
			deviceCnt--
			__this.authMutex.Unlock()

			log.Infof("Device %s disconnected. Remaining devices: %d\n", DeviceList.Device[i].Id, deviceCnt)
			return
		}
	}
}
