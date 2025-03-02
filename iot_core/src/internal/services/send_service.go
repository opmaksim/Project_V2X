package service

import (
	"encoding/json"
	config "project/v2x/iot-core/internal/configs"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SendProcessor struct {
	client mqtt.Client
}

var (
	lastInsert int64
	prevInsert int64
	stopFlag   int
	//mqttClient mqtt.Client
)

func (__this *SendProcessor) Connect(__broker_addr string, __topic_name string, __db *gorm.DB) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(__broker_addr)
	opts.SetClientID(config.MQTT_CLIENT_ID)

	__this.client = mqtt.NewClient(opts)
	if token := __this.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT 연결 실패: %v", token.Error())
	}
	log.Infof("Start MQTT Publisher: %s", __topic_name)

	// 주기적으로 DB 이벤트를 감지하도록 고루틴 사용
	go __this.monitorEvents(__topic_name, __db)
}

func (__this *SendProcessor) monitorEvents(__topic_name string, __db *gorm.DB) {
	// 주기적으로 이벤트 확인 (예: 10초마다)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := __this.checkForRecvEvent(__topic_name, __db)
			if err != nil {
				log.Error("이벤트 확인 실패")
			}
		}
	}
}

// MQTT를 통해 플래그 값 전송
func (__this *SendProcessor) messageTransmit(__db *gorm.DB, __topic_name string, __flag int) {
	if __this.client == nil || !__this.client.IsConnected() {
		log.Warn("MQTT 클라이언트가 연결되지 않음. 메시지 전송 불가.")
		return
	}

	//topic := "flag/update"
	//message := __flag
	message := map[string]interface{}{
		"stop": __flag,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Errorf("Error marshalling message: %v", err)
		return
	}

	//token := __this.client.Publish(__topic_name, 0, false, message)
	token := __this.client.Publish(__topic_name, 0, false, messageJSON)
	token.Wait()
	log.Printf("MQTT로 전송됨 → 토픽: %s, 메시지: %v", __topic_name, string(messageJSON))

	err = __this.clearLogDB(__db)
	if err != nil {
		log.Errorf("Error deleting all data from Flag_Logs: %v", err)
	}

	stopFlag = 0
}

func (__this *SendProcessor) clearLogDB(__db *gorm.DB) error {
	deleteQuery := `DELETE FROM Flag_Logs`
	err := __db.Exec(deleteQuery).Error
	if err != nil {
		return err
	}

	log.Info("All data from Flag_Logs has been deleted.")
	return nil
}

// DB에서 최근 INSERT 확인 후 MQTT로 플래그 전송
func (__this *SendProcessor) checkForRecvEvent(__topic_name string, __db *gorm.DB) error {
	// 최신 두 개의 created_at 및 db 값 가져오기
	query := `SELECT created_at, db FROM Flag_Logs ORDER BY created_at DESC LIMIT 2`
	rows, err := __db.Raw(query).Rows()
	if err != nil {
		log.Errorf("Error querying to select: %v", err)
		return err
	}
	defer rows.Close()

	var latestTimestamp, prevTimestamp int64
	var latestDB, prevDB string

	// 첫 번째 레코드 읽기 (최신 데이터)
	if rows.Next() {
		var latestTime time.Time
		err = rows.Scan(&latestTime, &latestDB)
		if err != nil {
			log.Errorf("Error scanning latest timestamp and db: %v", err)
			return err
		}
		// time.Time을 Unix 타임스탬프(초)로 변환하여 int64로 저장
		latestTimestamp = latestTime.Unix()
	}

	// 두 번째 레코드 읽기 (이전 데이터)
	if rows.Next() {
		var prevTime time.Time
		err = rows.Scan(&prevTime, &prevDB)
		if err != nil {
			log.Errorf("Error scanning previous timestamp and db: %v", err)
			return err
		}
		// time.Time을 Unix 타임스탬프(초)로 변환하여 int64로 저장
		prevTimestamp = prevTime.Unix()
	}

	// 두 개의 db 값이 다를 경우 처리
	if latestTimestamp > lastInsert {
		lastInsert = latestTimestamp
		prevInsert = prevTimestamp

		if latestDB != prevDB {
			if stopFlag == 0 {
				stopFlag = 1
				go __this.messageTransmit(__db, __topic_name, stopFlag)
			}
		} else {
			log.Warn("DB is equal")
		}
	}

	return nil
}
