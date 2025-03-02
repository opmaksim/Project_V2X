#include <iostream>
#include <thread>
#include <termios.h>

#include <ros/ros.h>
#include <std_msgs/Int16.h>
#include <sstream>
#include "project_ros/project_ros_msg.h"
int getch(void)
{
  int ch;
  struct termios oldt;
  struct termios newt;

  // Store old settings, and copy to new settings
  tcgetattr(STDIN_FILENO, &oldt);
  newt = oldt;

  // Make required changes and apply the settings
  newt.c_lflag &= ~(ICANON | ECHO);
  tcsetattr(fileno(stdin), TCSANOW, &newt);
  // Get the current character
  ch = getchar();

  // Reapply old settings
  tcsetattr(STDIN_FILENO, TCSANOW, &oldt);

  return ch;
}
class RosDriver
{
private:
   ros::NodeHandle nh_;
   ros::Publisher pub;
public:
   double dir, spd;
   char cmd[50]; 
   RosDriver(ros::NodeHandle &nh)
   {
	nh_ = nh;
	dir =0, spd = 100;
	pub = nh_.advertise<project_ros::project_ros_msg>("motor_command",10);
   }

  bool driveKeyboard()
  {
    std::cout << "Type a command and then press enter.  "
      "Use 'w' to move forward, 'a' to turn left, "
      "'d' to turn right, 's' to stop, "
      "'x' to move back, '.' to exit.\n";


    while(nh_.ok()){

      //std::cin.getline(cmd, 50);
      cmd[0] = getch();
      if(cmd[0]!='w' && cmd[0]!='a' && cmd[0]!='d' && cmd[0]!='s' && cmd[0]!='x' && cmd[0]!='.')
      {
        std::cout << "unknown command:" << cmd << "\n";
        continue;
      }
      //move forward
      if(cmd[0]=='w'){
	if(spd < 250){spd+=50;}
	else{spd = 100;}
	dir = 1;
        std::cout<<"dir : "<<"foward"<<"\n";
        std::cout<<"spd : "<<spd<<"\n";
      }
      //turn left (yaw) and drive forward at the same time
      else if(cmd[0]=='a'){
	if(spd < 250)spd+=50;
	else{spd = 100;}
        dir = 3;
        std::cout<<"dir : "<<"left"<<"\n";
        std::cout<<"spd : "<<spd<<"\n";
      }
      //turn right (yaw) and drive forward at the same time
      else if(cmd[0]=='d'){
        if(spd < 250)spd+=50;
	else{spd = 100;}
        dir = 4;
        std::cout<<"dir : "<<"right"<<"\n";
        std::cout<<"spd : "<<spd<<"\n";
      }
      else if(cmd[0]=='s'){
        if(spd < 250)spd+=50;
	else{spd = 100;}
        dir = 2;
        std::cout<<"dir : "<<"back"<<"\n";
        std::cout<<"spd : "<<spd<<"\n";
      }
      //quit
      else if(cmd[0]=='.'){
        dir = 0;
        spd = 0;
	std::cout<<"dir : "<<"back"<<"\n";
        std::cout<<"spd : "<<spd<<"\n";

        break;
      }
    }
    return true;
  }

  bool send_msg_thread()
  {
    project_ros::project_ros_msg base_cmd;

    ros::Rate loop_rate(3);

    base_cmd.dir = 0;
    base_cmd.spd = 100;
    while(1)
    {
      //publish the assembled command
      base_cmd.dir = dir;
      base_cmd.spd = spd;
      pub.publish(base_cmd);
      if(cmd[0]=='.')break;

      loop_rate.sleep();
    }
    return true;
  }
  void robot_driver_start()
  {
    std::thread t1(&RosDriver::driveKeyboard, this);
    std::thread t2(&RosDriver::send_msg_thread, this);
    t1.join();
    t2.join();
  }
};




int main(int argc, char **argv)
{
  ros::init(argc, argv, "project_ros_pub");
  ros::NodeHandle nh;

  RosDriver driver(nh);
  driver.robot_driver_start();
}
