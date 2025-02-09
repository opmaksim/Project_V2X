#ifndef __MOTOR_H
#define __MOTOR_H

#include "stm32f4xx_hal.h"
#include "main.h"

extern TIM_HandleTypeDef htim2;

void motorForward();
void motorBack();
void motorSpeed(uint16_t motorPwm);
#endif