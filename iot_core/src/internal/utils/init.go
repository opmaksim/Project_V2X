package utils

import (
	auth "project/v2x/iot-core/internal/auths"

	log "github.com/sirupsen/logrus"
)

func init() {
	// .env 파일을 로드하여 환경 변수 설정
	//config.LoadEnv()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	auth.DeviceList = auth.DeviceAuth{
		Device: [auth.DEVICE_MAX_CLNT]auth.DeviceAccount{
			{Index: 0, Conn: nil, Ip: nil, Id: "PJW_ARD", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "PJW_TEST", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "SCM_CAR", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "4", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "5", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "6", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "7", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "8", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "9", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "10", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "11", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "12", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "13", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "14", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "15", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "16", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "17", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "18", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "19", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "20", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "21", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "22", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "23", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "24", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "25", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "26", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "27", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "28", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "29", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "30", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "31", Pw: "PASSWD"},
			{Index: 0, Conn: nil, Ip: nil, Id: "32", Pw: "PASSWD"},
		},
	}
}
