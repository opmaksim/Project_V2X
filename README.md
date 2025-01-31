# Project_V2X

## ~25.01.27
### Hardware Setting
<img src="img/01.hardware.png" alt="alt text" style="width: 30%;"/>


## 25.01.28

### function
- w,a,s,d 키보드로 car 제어
- . 입력시 정지

### setting
- [code](https://github.com/opmaksim/Project_V2X/tree/feature/ros_car/src/25.01.28_control_car_by_keyboard)
 1. 📂project_ros : jetsonNano ~/catkin_ws/src 에 붙여넣기
 2. 📂project_ros_arduino : Arduino IDE를 사용하여 OpenCR에 업로드
 3. ```bash
    #ROS host 시작
    roscore
 4. ```bash
    #Jetson과 OpenCr serial 통신
    rosrun rosserial_python serial_node.py __name:=arduino _port:=/dev/ttyACM0 _baud:=57600
 5. ```bash
    #publisher 실행
    rosrun project_ros project_ros

### challenge
- spd 변수 공유로인한 불편함 해결
- serial 통신 문제 원인 찾아야함 ( 전력 or wifi 문제 )

## 25.01.31

### function
- line follow </br>
![alt text](img/2.line_follow.gif)


### setting
- [code]
- OpenCR 보드엔 01.28 코드 그대로 업로드
- requirement.txt 설치 후 코드 실행

###challenge
- 가끔 opencr이 동작 안하는 문제 (바퀴를 손으로 돌려주다보면 갑자기 동작함 왜 이런지는 모르겠음)
- 두개의 라인을 따라가도록 변경 ( 실제 도로와 유사환경 )
