#ifndef __HC_SR04_H
#define __HC_SR04_H

#include "stm32f4xx_hal.h"
#include "main.h"

#define HCSR04_TIMEOUT              1000000
#define TRIG_PIN_1_Pin              GPIO_PIN_11
#define TRIG_PIN_1_GPIO_Port        GPIOA
#define ECHO_PIN_1_Pin              GPIO_PIN_12
#define ECHO_PIN_1_GPIO_Port        GPIOB
#define TRIG_PIN_2_Pin              GPIO_PIN_2
#define TRIG_PIN_2_GPIO_Port        GPIOB
#define ECHO_PIN_2_Pin              GPIO_PIN_1
#define ECHO_PIN_2_GPIO_Port        GPIOB
#define TRIG_PIN_3_Pin              GPIO_PIN_13
#define TRIG_PIN_3_GPIO_Port        GPIOB
#define ECHO_PIN_3_Pin              GPIO_PIN_4
#define ECHO_PIN_3_GPIO_Port        GPIOC


extern volatile float distance[4];  // 각 센서의 거리 저장 (소숫점 포함)
extern TIM_HandleTypeDef htim1;
void _delay_us(uint16_t us);
void trigger_pulse(uint8_t sensor);
float measure_distance(uint8_t sensor);
#endif