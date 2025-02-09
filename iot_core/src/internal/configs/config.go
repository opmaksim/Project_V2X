package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

const (
	DB_HOST   = "10.10.14.34"
	DB_USER   = "v2x_user"
	DB_PASSWD = "pwv2x"
	DB_NAME   = "v2x"

	MQTT_BROKER_ADDRESS            = "tcp://10.10.14.34:1883"
	MQTT_DATA_EOTION_COLLECT_TOPIC = "project/v2x/iot-core/data-emotion-collector"
	MQTT_DATA_DRIVE_COLLECT_TOPIC  = "project/v2x/iot-core/data-drive-collector"
	MQTT_CAR_CONTROL_TOPIC         = "project/v2x/iot-core/car-controller"
	MQTT_CLIENT_ID                 = "IotCore"

	KUBE_SECRET_PATH      = "internal/tls/"
	KUBE_API_SERVER       = "https://10.10.14.34:6443"
	KUBE_CA_CERT_PATH     = KUBE_SECRET_PATH + "ca.crt"
	KUBE_CLIENT_CERT_PATH = KUBE_SECRET_PATH + "iot-core.crt"
	KUBE_CLIENT_KEY_PATH  = KUBE_SECRET_PATH + "iot-core.key"
	KUBE_CONFIG           = "internal/config/kube_config"
)

func LoadEnv() {
	// .env 파일을 로드하여 환경 변수 설정
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
