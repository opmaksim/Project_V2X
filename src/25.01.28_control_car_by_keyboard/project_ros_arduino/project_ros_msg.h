#ifndef _ROS_project_ros_project_ros_msg_h
#define _ROS_project_ros_project_ros_msg_h

#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "ros/msg.h"

namespace project_ros
{

  class project_ros_msg : public ros::Msg
  {
    public:
      typedef int16_t _dir_type;
      _dir_type dir;
      typedef int16_t _spd_type;
      _spd_type spd;

    project_ros_msg():
      dir(0),
      spd(0)
    {
    }

    virtual int serialize(unsigned char *outbuffer) const override
    {
      int offset = 0;
      union {
        int16_t real;
        uint16_t base;
      } u_dir;
      u_dir.real = this->dir;
      *(outbuffer + offset + 0) = (u_dir.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_dir.base >> (8 * 1)) & 0xFF;
      offset += sizeof(this->dir);
      union {
        int16_t real;
        uint16_t base;
      } u_spd;
      u_spd.real = this->spd;
      *(outbuffer + offset + 0) = (u_spd.base >> (8 * 0)) & 0xFF;
      *(outbuffer + offset + 1) = (u_spd.base >> (8 * 1)) & 0xFF;
      offset += sizeof(this->spd);
      return offset;
    }

    virtual int deserialize(unsigned char *inbuffer) override
    {
      int offset = 0;
      union {
        int16_t real;
        uint16_t base;
      } u_dir;
      u_dir.base = 0;
      u_dir.base |= ((uint16_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_dir.base |= ((uint16_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->dir = u_dir.real;
      offset += sizeof(this->dir);
      union {
        int16_t real;
        uint16_t base;
      } u_spd;
      u_spd.base = 0;
      u_spd.base |= ((uint16_t) (*(inbuffer + offset + 0))) << (8 * 0);
      u_spd.base |= ((uint16_t) (*(inbuffer + offset + 1))) << (8 * 1);
      this->spd = u_spd.real;
      offset += sizeof(this->spd);
     return offset;
    }

    virtual const char * getType() override { return "project_ros/project_ros_msg"; };
    virtual const char * getMD5() override { return "1ebfd9092e958b03909c95e1382c65a6"; };

  };

}
#endif
