#include <ros.h>
#include <std_msgs/Int16.h>
#include "project_ros_msg.h"

// Pin 설정
int Dir1Pin_A = 2; // 제어 신호 1핀
int Dir2Pin_A = 3; // 제어 신호 2핀
int Dir1Pin_B = 4;
int Dir2Pin_B = 5;
int SpeedPin_A = 10; // PWM 제어를 위한 핀
int SpeedPin_B = 11; // PWM 제어를 위한 핀

class MotorController {
public:
    double spd;
    MotorController() : spd(255) {}
    void forward() {
        digitalWrite(Dir1Pin_A, HIGH);
        digitalWrite(Dir2Pin_A, LOW);
        digitalWrite(Dir1Pin_B, HIGH);
        digitalWrite(Dir2Pin_B, LOW);
        analogWrite(SpeedPin_A, spd);
        analogWrite(SpeedPin_B, spd);
    }

    void backward() {
        digitalWrite(Dir1Pin_A, LOW);
        digitalWrite(Dir2Pin_A, HIGH);
        digitalWrite(Dir1Pin_B, LOW);
        digitalWrite(Dir2Pin_B, HIGH);
        analogWrite(SpeedPin_A, spd);
        analogWrite(SpeedPin_B, spd);
    }

    void turnLeft() {
        digitalWrite(Dir1Pin_A, LOW);
        digitalWrite(Dir2Pin_A, HIGH);
        digitalWrite(Dir1Pin_B, HIGH);
        digitalWrite(Dir2Pin_B, LOW);
        analogWrite(SpeedPin_A, spd * 0.6);
        analogWrite(SpeedPin_B, spd);
    }

    void turnRight() {
        digitalWrite(Dir1Pin_A, HIGH);
        digitalWrite(Dir2Pin_A, LOW);
        digitalWrite(Dir1Pin_B, LOW);
        digitalWrite(Dir2Pin_B, HIGH);
        analogWrite(SpeedPin_A, spd);
        analogWrite(SpeedPin_B, spd * 0.6);
    }
};

ros::NodeHandle nh;
MotorController motor;


void commandCallback(const project_ros::project_ros_msg& cmd_msg) {
    motor.spd = cmd_msg.spd;
    if (cmd_msg.dir == 1) {
        motor.forward();
    } else if (cmd_msg.dir == 2) {
        motor.backward();
    } else if (cmd_msg.dir == 3) {
        motor.turnLeft();
    } else if (cmd_msg.dir == 4) {
        motor.turnRight();
    }
}

ros::Subscriber<project_ros::project_ros_msg> sub("motor_command", &commandCallback);

void setup() {
    pinMode(Dir1Pin_A, OUTPUT);
    pinMode(Dir2Pin_A, OUTPUT);
    pinMode(Dir1Pin_B, OUTPUT);
    pinMode(Dir2Pin_B, OUTPUT);
    pinMode(SpeedPin_A, OUTPUT);
    pinMode(SpeedPin_B, OUTPUT);
    nh.initNode();
    nh.subscribe(sub);
}

void loop() {
    nh.spinOnce();
    delay(10);
}
