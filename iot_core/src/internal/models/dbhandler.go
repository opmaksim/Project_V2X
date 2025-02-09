package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

//======================================================================================= DB: Device_Info =======================================================================================//

func (__this *DeviceInfo) GetTables(__db *gorm.DB) ([]interface{}, error) {
	query := `SELECT device_id, device_type, device_name, connected_status, drive_status, connected_time FROM Device_Info`
	rows, err := __db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []interface{}
	for rows.Next() {
		var data DeviceInfo
		if err := rows.Scan(&data.DeviceId, &data.DeviceType, &data.DeviceName, &data.ConnectedStatus, &data.DriveStatus, &data.ConnetedTime); err != nil {
			return nil, err
		}

		events = append(events, data)
	}

	return events, nil
}

func (__this *DeviceInfo) InsertTables(__db *gorm.DB) error {
	query := `INSERT INTO Device_Info (device_id, device_type, device_name, connected_status, drive_status, connected_time) VALUES (?, ?, ?, ?, ?, ?)`
	res := __db.Exec(query, __this.DeviceId, __this.DeviceType, __this.DeviceName, __this.ConnectedStatus, __this.DriveStatus, __this.ConnetedTime)
	if res.Error != nil {
		log.Errorf("Error inserting into Device_Info: %v", res.Error)
		return res.Error
	}
	fmt.Println("Device_Info inserted successfully")

	return nil
}

func (__this *DeviceInfo) UpdateTables(__db *gorm.DB) error {
	query := `UPDATE Device_Info SET device_type = ?, device_name = ?, connected_status = ?, drive_status = ?, connected_time = ? WHERE device_id = ?`
	res := __db.Exec(query, __this.DeviceType, __this.DeviceName, __this.ConnectedStatus, __this.DriveStatus, __this.ConnetedTime, __this.DeviceId)
	if res.Error != nil {
		log.Errorf("Error inserting updating Device_Info: %v", res.Error)
		return res.Error
	}
	log.Info("Device_Info updated successfully")

	return nil
}

func (__this *DeviceInfo) DeleteTables(__db *gorm.DB) error {
	query := `DELETE FROM Device_Info WHERE device_id = ?`
	res := __db.Exec(query, __this.DeviceId)
	if res.Error != nil {
		log.Errorf("Error deleting into Device_Info: %v", res.Error)
		return res.Error
	}
	log.Info("Device_Info deleted successfully")

	return nil
}

//======================================================================================= DB: Accident_Status =======================================================================================//

func (__this *AccidentStatus) GetTables(__db *gorm.DB) ([]interface{}, error) {
	query := `SELECT device_id, emotion, event_time FROM Accident_Status`
	rows, err := __db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []interface{}
	for rows.Next() {
		var data AccidentStatus
		// if err := rows.Scan(&data.DeviceId, &data.Drowsiness, &data.Unintended, &data.Sudden, &data.EventTime); err != nil {
		// 	return nil, err
		// }
		if err := rows.Scan(&data.DeviceId, &data.Emotion, &data.EventTime); err != nil {
			return nil, err
		}

		events = append(events, data)
	}

	return events, nil
}

func (__this *AccidentStatus) InsertTables(__db *gorm.DB) error {
	query := `INSERT INTO Accident_Status (device_id, emotion, event_time) VALUES (?, ?, ?)`
	//res := __db.Exec(query, __this.DeviceId, __this.Drowsiness, __this.Unintended, __this.Sudden, __this.EventTime)
	res := __db.Exec(query, __this.DeviceId, __this.Emotion, __this.EventTime)
	if res.Error != nil {
		log.Errorf("Error inserting into Accident_Status: %v", res.Error)
		return res.Error
	}
	log.Info("Accident_Status inserted successfully")

	return nil
}

func (__this *AccidentStatus) UpdateTables(__db *gorm.DB) error {
	query := `UPDATE Accident_Status SET emotion = ? WHERE device_id = ?`
	//res := __db.Exec(query, __this.Drowsiness, __this.Unintended, __this.Sudden, __this.EventTime, __this.DeviceId)
	res := __db.Exec(query, __this.Emotion, __this.EventTime, __this.DeviceId)
	if res.Error != nil {
		log.Errorf("Error inserting updating Accident_Status: %v", res.Error)
		return res.Error
	}
	log.Info("Accident_Status updated successfully")

	return nil
}

func (__this *AccidentStatus) DeleteTables(__db *gorm.DB) error {
	query := `DELETE FROM Accident_Status WHERE device_id = ?`
	res := __db.Exec(query, __this.DeviceId)
	if res.Error != nil {
		log.Errorf("Error deleting into Accident_Status: %v", res.Error)
		return res.Error
	}
	log.Info("Accident_Status deleted successfully")

	return nil
}

//======================================================================================= DB: Drive_Events =======================================================================================//

func (__this *DriveEvents) GetTables(__db *gorm.DB) ([]interface{}, error) {
	query := `SELECT device_id, handle, brake, accel, pressure, drive_time FROM Drive_Events`
	rows, err := __db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []interface{}
	for rows.Next() {
		var data DriveEvents
		if err := rows.Scan(&data.DeviceId, &data.Handle, &data.Brake, &data.Accel, &data.Pressure, &data.DriveTime); err != nil {
			return nil, err
		}

		events = append(events, data)
	}

	return events, nil
}

func (__this *DriveEvents) InsertTables(__db *gorm.DB) error {
	query := `INSERT INTO Drive_Events (device_id, handle, brake, accel, pressure, drive_time) VALUES (?, ?, ?, ?, ?, ?)`
	res := __db.Exec(query, __this.DeviceId, __this.Handle, __this.Brake, __this.Accel, __this.Pressure, __this.DriveTime)
	if res.Error != nil {
		log.Errorf("Error inserting into Drive_Events: %v", res.Error)
		return res.Error
	}
	log.Info("Drive_Events inserted successfully")

	return nil
}

func (__this *DriveEvents) UpdateTables(__db *gorm.DB) error {
	query := `UPDATE Drive_Events SET handle = ?, brake = ?, accel = ?, drive_time = ?, pressure = ? WHERE device_id = ?`
	res := __db.Exec(query, __this.Handle, __this.Brake, __this.Accel, __this.DriveTime, __this.Pressure, __this.DeviceId)
	if res.Error != nil {
		log.Errorf("Error inserting updating Drive_Events: %v", res.Error)
		return res.Error
	}
	log.Info("Drive_Events updated successfully")

	return nil
}

func (__this *DriveEvents) DeleteTables(__db *gorm.DB) error {
	query := `DELETE FROM Drive_Events WHERE device_id = ?`
	res := __db.Exec(query, __this.DeviceId)
	if res.Error != nil {
		log.Errorf("Error deleting into Drive_Events: %v", res.Error)
		return res.Error
	}
	log.Info("Drive_Events deleted successfully")

	return nil
}
