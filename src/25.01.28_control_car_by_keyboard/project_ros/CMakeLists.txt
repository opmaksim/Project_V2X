cmake_minimum_required(VERSION 2.8.3)
project(project_ros)
find_package(catkin REQUIRED COMPONENTS roscpp std_msgs message_generation)
add_message_files(FILES project_ros_msg.msg)
generate_messages(
   DEPENDENCIES
   std_msgs
 )
catkin_package(CATKIN_DEPENDS roscpp std_msgs)
include_directories(${catkin_INCLUDE_DIRS})
add_executable(project_ros src/project_ros_pub.cpp)
target_link_libraries(project_ros ${catkin_LIBRARIES})
