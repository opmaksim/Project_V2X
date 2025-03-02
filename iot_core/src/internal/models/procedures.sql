-- 기존 프로시저가 있으면 삭제
DROP PROCEDURE IF EXISTS EnsureConstraints;

DELIMITER $$

CREATE PROCEDURE `EnsureConstraints`()
BEGIN
    -- 'Device_Info' 테이블의 기본 키 추가
    IF NOT EXISTS (
        SELECT 1
        FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS
        WHERE CONSTRAINT_TYPE = 'PRIMARY KEY'
        AND TABLE_NAME = 'Device_Info'
        AND TABLE_SCHEMA = 'v2x'
    ) THEN
        ALTER TABLE `Device_Info`
        ADD CONSTRAINT `PK_DEVICE_INFO` PRIMARY KEY (`device_id`);
    END IF;

    -- 'Drive_Events' 테이블의 기본 키 추가
    IF NOT EXISTS (
        SELECT 1
        FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS
        WHERE CONSTRAINT_TYPE = 'PRIMARY KEY'
        AND TABLE_NAME = 'Drive_Events'
        AND TABLE_SCHEMA = 'v2x'
    ) THEN
        ALTER TABLE `Drive_Events`
        ADD CONSTRAINT `PK_DRIVE_EVENTS` PRIMARY KEY (`seq`, `device_id`);
    END IF;

    -- 'Accident_Status' 테이블의 기본 키 추가
    IF NOT EXISTS (
        SELECT 1
        FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS
        WHERE CONSTRAINT_TYPE = 'PRIMARY KEY'
        AND TABLE_NAME = 'Accident_Status'
        AND TABLE_SCHEMA = 'v2x'
    ) THEN
        ALTER TABLE `Accident_Status`
        ADD CONSTRAINT `PK_ACCIDENT_STATUS` PRIMARY KEY (`seq`, `device_id`);
    END IF;

    -- 'Drive_Events' 테이블의 외래 키 추가
    IF NOT EXISTS (
        SELECT 1
        FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE
        WHERE CONSTRAINT_NAME = 'FK_Device_Info_TO_Drive_Events_1'
        AND TABLE_NAME = 'Drive_Events'
        AND TABLE_SCHEMA = 'v2x'
    ) THEN
        ALTER TABLE `Drive_Events`
        ADD CONSTRAINT `FK_Device_Info_TO_Drive_Events_1`
        FOREIGN KEY (`device_id`) REFERENCES `Device_Info` (`device_id`);
    END IF;

    -- 'Accident_Status' 테이블의 외래 키 추가
    IF NOT EXISTS (
        SELECT 1
        FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE
        WHERE CONSTRAINT_NAME = 'FK_Device_Info_TO_Accident_Status_1'
        AND TABLE_NAME = 'Accident_Status'
        AND TABLE_SCHEMA = 'v2x'
    ) THEN
        ALTER TABLE `Accident_Status`
        ADD CONSTRAINT `FK_Device_Info_TO_Accident_Status_1`
        FOREIGN KEY (`device_id`) REFERENCES `Device_Info` (`device_id`);
    END IF;

END$$

DELIMITER ;