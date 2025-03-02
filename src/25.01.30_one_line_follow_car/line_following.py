import cv2
import numpy as np
import rospy
from project_ros.msg import project_ros_msg

# Initialize ROS node
rospy.init_node('line_follower')
pub = rospy.Publisher('motor_command', project_ros_msg, queue_size=10)

# Function to publish motor commands
def publish_motor_command(direction, speed):
    msg = project_ros_msg()
    msg.dir = direction
    msg.spd = speed
    pub.publish(msg)

cap = cv2.VideoCapture(0)
cap.set(3, 160)
cap.set(4, 120)

while not rospy.is_shutdown():
    ret, frame = cap.read()
    low_b = np.uint8([5, 5, 5])
    high_b = np.uint8([0, 0, 0])
    mask = cv2.inRange(frame, high_b, low_b)
    contours, hierarchy = cv2.findContours(mask, 1, cv2.CHAIN_APPROX_NONE)

    if len(contours) > 0:
        con = max(contours, key=cv2.contourArea)
        M = cv2.moments(con)
        if M["m00"] != 0:
            cx = int(M['m10'] / M['m00'])
            cy = int(M['m01'] / M['m00'])
            print("CX :" + str(cx) + " CY:" + str(cy))

            if cx >= 115:
                print("Turn Right")
                publish_motor_command(4, 165)  # Turn Left
            elif cx >= 45 and cx < 115:
                print("On Track")
                publish_motor_command(1, 140)  # Move Forward
            elif cx <= 45:
                print("Turn Left")
                publish_motor_command(3, 165)  # Turn Right

            cv2.circle(frame, (cx, cy), 5, (255, 255, 255), -1)

    cv2.drawContours(frame, con, -1, (0, 255, 0), 1)


cap.release()
cv2.destroyAllWindows()