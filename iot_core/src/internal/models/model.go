package model

import (
	"fmt"
	"log"
	_ "os"
	config "project/v2x/iot-core/internal/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnector struct {
	// priavate
	dbHost   string
	dbUser   string
	dbPasswd string
	dbName   string
	// driveEventsTrigger    string
	// accidentStatusTrigger string
	//triggers string

	DB *gorm.DB
}

// DB: Device_Info
type DeviceInfo struct {
	// public
	DeviceId        string `json:"device_id"`
	DeviceType      string `json:"device_type"`
	DeviceName      string `json:"device_name"`
	ConnectedStatus bool   `json:"connected_status"`
	DriveStatus     bool   `json:"drive_status"`
	ConnetedTime    string `json:"connected_time"`
}

// DB: Accident_Status
type AccidentStatus struct {
	// public
	DeviceId string `json:"device_id"`
	// Drowsiness int    `json:"drowsiness"`
	// Unintended int    `json:"unintended"`
	// Sudden     int    `json:"sudden"`
	Emotion   bool   `json:"emotion"`
	EventTime string `json:"event_time"`
}

// DB: Drive_Events
type DriveEvents struct {
	// public
	DeviceId  string `json:"device_id"`
	Handle    int    `json:"handle"`
	Brake     int    `json:"brake"`
	Accel     int    `json:"accel"`
	Pressure  int    `json:"pressure"`
	DriveTime string `json:"drive_time"`
}

// DB: Flag_Logs
type FlagLogs struct {
	Seq       uint   `gorm:"primaryKey"`
	Message   string `gorm:"size:255"`
	DB        string `gorm:"size:40"`
	Sent      bool   `gorm:"default:false"`
	CreatedAt string `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

// DB 구조체의 메서드에 대한 인터페이스
type DBHandler interface {
	GetTables(__db *gorm.DB) ([]interface{}, error)
	InsertTables(__db *gorm.DB) error
	UpdateTables(__db *gorm.DB) error
	DeleteTables(__db *gorm.DB) (int64, error)
}

func InitDBInfo() *DBConnector {
	return &DBConnector{
		// dbHost:   os.Getenv("DB_HOST"),
		// dbUser:   os.Getenv("DB_USER"),
		// dbPasswd: os.Getenv("DB_PASSWD"),
		// dbName:   os.Getenv("DB_NAME"),
		// driveEventsTrigger:    os.Getenv("DRIVE_EVENTS_TRIGGER"),
		// accidentStatusTrigger: os.Getenv("ACCIDENTS_STATUS_TRIGGER"),
		//triggers: os.Getenv("TRIGGER_FILE"),
		dbHost:   config.DB_HOST,
		dbUser:   config.DB_USER,
		dbPasswd: config.DB_PASSWD,
		dbName:   config.DB_NAME,
	}
}

// DB에 연결하기 위한 메서드
func (__this *DBConnector) Connect() (*gorm.DB, error) {
	// MariaDB 연결
	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", __this.dbUser, __this.dbPasswd, __this.dbHost, __this.dbName)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 연결된 DB 객체 설정
	__this.DB = db

	return __this.DB, nil
}

// func (__this *DBConnector) executeSQLFile(__db *gorm.DB, __filePath string) error {
// 	// 파일을 읽어옵니다.
// 	sqlBytes, err := os.ReadFile(__filePath)
// 	if err != nil {
// 		return fmt.Errorf("failed to read SQL file %s: %v", __filePath, err)
// 	}

// 	// SQL 쿼리 실행
// 	sqlQuery := string(sqlBytes)
// 	if err := __db.Exec(sqlQuery).Error; err != nil {
// 		return fmt.Errorf("failed to execute SQL file %s: %v", __filePath, err)
// 	}

// 	return nil
// }

func (__this *DBConnector) CreateTables(__db *gorm.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS Device_Info (
			device_id VARCHAR(20) NOT NULL,
			device_type VARCHAR(20) NULL,
			device_name VARCHAR(20) NULL,
			connected_status BOOLEAN NULL,
			drive_status BOOLEAN NULL,
			connected_time TIMESTAMP NULL,
			PRIMARY KEY (device_id)
		);
	`
	if err := __db.Exec(query).Error; err != nil {
		log.Printf("Error creating Device_Info table: %v", err)
	} else {
		fmt.Println("Device_Info table created successfully!")
	}

	query = `
		CREATE TABLE IF NOT EXISTS Drive_Events (
			seq INT AUTO_INCREMENT NOT NULL,
			device_id VARCHAR(20) NOT NULL,
			handle INT NULL,
			brake INT NULL,
			accel INT NULL,
			pressure INT NULL,
			drive_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (seq),
			FOREIGN KEY (device_id) REFERENCES Device_Info (device_id) ON DELETE CASCADE
		);
	`
	if err := __db.Exec(query).Error; err != nil {
		log.Printf("Error creating Drive_Events table: %v", err)
	} else {
		fmt.Println("Drive_Events table created successfully!")
	}

	// query = `
	// 	CREATE TABLE IF NOT EXISTS Accident_Status (
	// 		seq INT AUTO_INCREMENT NOT NULL,
	// 		device_id VARCHAR(20) NOT NULL,
	// 		drowsiness BOOLEAN NULL,
	// 		unintended BOOLEAN NULL,
	// 		sudden BOOLEAN NULL,
	// 		event_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	// 		PRIMARY KEY (seq),
	// 		FOREIGN KEY (device_id) REFERENCES Device_Info (device_id) ON DELETE CASCADE
	// 	);
	// `
	query = `
		CREATE TABLE IF NOT EXISTS Accident_Status (
			seq INT AUTO_INCREMENT NOT NULL,
			device_id VARCHAR(20) NOT NULL,
			emotion BOOLEAN NULL,
			event_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (seq),
			FOREIGN KEY (device_id) REFERENCES Device_Info (device_id) ON DELETE CASCADE
		);
	`
	if err := __db.Exec(query).Error; err != nil {
		log.Printf("Error creating Accident_Status table: %v", err)
	} else {
		fmt.Println("Accident_Status table created successfully!")
	}

	query = `
		CREATE TABLE IF NOT EXISTS Flag_Logs (
			seq INT AUTO_INCREMENT NOT NULL,
			message VARCHAR(255),
			db VARCHAR(40),
			sent BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (seq)
		);
	`
	if err := __db.Exec(query).Error; err != nil {
		log.Printf("Error creating Flag_Logs table: %v", err)
	} else {
		fmt.Println("Flag_Logs table created successfully!")
	}

	// sqlFiles := []string{
	// 	// __this.accidentStatusTrigger,
	// 	// __this.driveEventsTrigger,
	// 	__this.triggers,
	// }

	// for _, filePath := range sqlFiles {
	// 	if err := __this.executeSQLFile(__db, filePath); err != nil {
	// 		log.Printf("Error executing SQL file %s: %v", filePath, err)
	// 	} else {
	// 		fmt.Printf("Successfully executed %s\n", filePath)
	// 	}
	// }
}

// DB 연결 종료를 위한 메서드
func (__this *DBConnector) Disconnect() {
	if __this.DB != nil {
		// gorm.DB에서 내부 *sql.DB 객체를 가져옴
		db, err := __this.DB.DB()
		if err != nil {
			log.Println("Error getting DB connection: ", err)
			return
		}

		// DB 연결 종료
		err = db.Close()
		if err != nil {
			log.Println("Error closing DB connection: ", err)
		} else {
			fmt.Println("DB connection closed successfully")
		}
	}
}
