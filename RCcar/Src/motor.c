#include "motor.h"

void motorForward(){
    HAL_GPIO_WritePin(MOTOR_PIN1_GPIO_Port, MOTOR_PIN1_Pin, RESET);
    HAL_GPIO_WritePin(MOTOR_PIN2_GPIO_Port, MOTOR_PIN2_Pin, SET);
}

void motorBack(){
    HAL_GPIO_WritePin(MOTOR_PIN1_GPIO_Port, MOTOR_PIN1_Pin, SET);
    HAL_GPIO_WritePin(MOTOR_PIN2_GPIO_Port, MOTOR_PIN2_Pin, RESET);
}

void motorSpeed(uint16_t motorPwm){
    TIM2->CCR1 = motorPwm;
}
