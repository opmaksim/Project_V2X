# Project_V2X

## ~25.01.27
### Hardware Setting
***
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
***
