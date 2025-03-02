-- 구분자 변경
DELIMITER $$

-- Accident_Status 테이블에 대한 트리거
DROP TRIGGER IF EXISTS after_accident_status_insert;
CREATE TRIGGER after_accident_status_insert
AFTER INSERT ON Accident_Status
FOR EACH ROW
BEGIN
    INSERT INTO Flag_Logs (message, db) 
    VALUES (CONCAT('New Accident Status inserted: device_id=', NEW.device_id, ', seq=', NEW.seq), 'Accident_Status');
END$$

-- Drive_Events 테이블에 대한 트리거
DROP TRIGGER IF EXISTS after_drive_events_insert;
CREATE TRIGGER after_drive_events_insert
AFTER INSERT ON Drive_Events
FOR EACH ROW
BEGIN
    INSERT INTO Flag_Logs (message, db) 
    VALUES (CONCAT('New Drive Event inserted: device_id=', NEW.device_id, ', seq=', NEW.seq), 'Drive_Events');
END$$

-- 구분자 기본값으로 변경
DELIMITER ;
