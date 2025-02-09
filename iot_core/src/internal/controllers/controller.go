package controller

import (
	config "project/v2x/iot-core/internal/configs"
	service "project/v2x/iot-core/internal/services"
	_ "sync"

	"gorm.io/gorm"
)

type AccidentStatusHandler struct{}
type DriveEventsHandler struct{}
type SendHandler struct{}

// 사고 상태 핸들러
func (__this *AccidentStatusHandler) Handler(__db *gorm.DB) {
	accChannel := &service.AccidentStatusProcessor{}

	go accChannel.Connect(config.MQTT_BROKER_ADDRESS, config.MQTT_DATA_EOTION_COLLECT_TOPIC, __db)
}

// 주행 이벤트 핸들러
func (__this *DriveEventsHandler) Handler(__db *gorm.DB) {
	driveChannel := &service.DriveEventsProcessor{}

	//go driveChannel.Connect(__db)
	go driveChannel.Connect(config.MQTT_BROKER_ADDRESS, config.MQTT_DATA_DRIVE_COLLECT_TOPIC, __db)
}

// MQTT Publisher 실행
func (__this *SendHandler) Handler(__db *gorm.DB) {
	sendChannel := &service.SendProcessor{}

	go sendChannel.Connect(config.MQTT_BROKER_ADDRESS, config.MQTT_CAR_CONTROL_TOPIC, __db)
}
