#include "HC_SR04.h"


void _delay_us(uint16_t us) {
    __HAL_TIM_SET_COUNTER(&htim1, 0); // 타이머 카운터 초기화
    while (__HAL_TIM_GET_COUNTER(&htim1) < us); // 원하는 시간만큼 대기
}

void trigger_pulse(uint8_t sensor) {
    // 각 센서에 따라 다른 핀 사용
    if (sensor == 0) {
        HAL_GPIO_WritePin(TRIG_PIN_1_GPIO_Port, TRIG_PIN_1_Pin, GPIO_PIN_RESET);
        _delay_us(2);
        HAL_GPIO_WritePin(TRIG_PIN_1_GPIO_Port, TRIG_PIN_1_Pin, GPIO_PIN_SET);
        _delay_us(10);
        HAL_GPIO_WritePin(TRIG_PIN_1_GPIO_Port, TRIG_PIN_1_Pin, GPIO_PIN_RESET);
    }
    else if (sensor == 1) {
        HAL_GPIO_WritePin(TRIG_PIN_2_GPIO_Port, TRIG_PIN_2_Pin, GPIO_PIN_RESET);
        _delay_us(2);
        HAL_GPIO_WritePin(TRIG_PIN_2_GPIO_Port, TRIG_PIN_2_Pin, GPIO_PIN_SET);
        _delay_us(10);
        HAL_GPIO_WritePin(TRIG_PIN_2_GPIO_Port, TRIG_PIN_2_Pin, GPIO_PIN_RESET);
    }
    else if (sensor == 2) {
        HAL_GPIO_WritePin(TRIG_PIN_3_GPIO_Port, TRIG_PIN_3_Pin, GPIO_PIN_RESET);
        _delay_us(2);
        HAL_GPIO_WritePin(TRIG_PIN_3_GPIO_Port, TRIG_PIN_3_Pin, GPIO_PIN_SET);
        _delay_us(10);
        HAL_GPIO_WritePin(TRIG_PIN_3_GPIO_Port, TRIG_PIN_3_Pin, GPIO_PIN_RESET);
    }
}

float measure_distance(uint8_t sensor) {
    uint32_t time = 0;
    uint32_t timeout = HCSR04_TIMEOUT;

    trigger_pulse(sensor);

    // 에코 핀 읽기
    if (sensor == 0) {
        while (HAL_GPIO_ReadPin(ECHO_PIN_1_GPIO_Port, ECHO_PIN_1_Pin) == GPIO_PIN_RESET) {
            if(timeout-- == 0) return -1;
        }
        while (HAL_GPIO_ReadPin(ECHO_PIN_1_GPIO_Port, ECHO_PIN_1_Pin) == GPIO_PIN_SET) {
            if(time++ >= 5000) break;
            _delay_us(1);
        }
    }
    else if (sensor == 1) {
        while (HAL_GPIO_ReadPin(ECHO_PIN_2_GPIO_Port, ECHO_PIN_2_Pin) == GPIO_PIN_RESET) {
            if(timeout-- == 0) return -1;
        }
        while (HAL_GPIO_ReadPin(ECHO_PIN_2_GPIO_Port, ECHO_PIN_2_Pin) == GPIO_PIN_SET) {
            if(time++ >= 5000) break;
            _delay_us(1);
        }
    }
    else if (sensor == 2) {
        while (HAL_GPIO_ReadPin(ECHO_PIN_3_GPIO_Port, ECHO_PIN_3_Pin) == GPIO_PIN_RESET) {
            if(timeout-- == 0) return -1;
        }
        while (HAL_GPIO_ReadPin(ECHO_PIN_3_GPIO_Port, ECHO_PIN_3_Pin) == GPIO_PIN_SET) {
            if(time++ >= 5000) break;
            _delay_us(1);
        }
    }

    // 거리 계산 (소수점 포함)
    float distance = (time * 0.0343) / 2; // 0.0343 cm/μs를 사용하여 거리 계산
    return distance;
}
