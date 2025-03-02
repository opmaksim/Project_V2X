#include <iostream>
#include <mqtt/async_client.h>
#include <cstdlib>
#include <chrono>
#include <ctime>

#define ADDRESS "tcp://10.10.14.34:1883" // MQTT 브로커 주소
#define CLIENT_ID "IotCorePublisher"
#define TOPIC "v2x/iot-core/accident-status"

using namespace std;

int main() {
    srand(time(0)); // 랜덤 값을 위한 초기화

    mqtt::async_client client(ADDRESS, CLIENT_ID);

    // 연결 옵션 설정
    mqtt::connect_options connOpts;
    connOpts.set_clean_session(true);

    string device_id = "car_01";
    int seq = 0;

    try {
        // 브로커 연결
        cout << "Connecting to the MQTT broker at " << ADDRESS << "..." << endl;
        client.connect(connOpts)->wait();
        cout << "Connected successfully." << endl;

        // 데이터 전송
        while (true) {
            // int handle = 1300 + static_cast<int>(rand()) / (static_cast<int>(RAND_MAX / 100));
            // int stop = 1300 + static_cast<int>(rand()) / (static_cast<int>(RAND_MAX / 200));
            // int accel = 1300 + static_cast<int>(rand()) / (static_cast<int>(RAND_MAX / 300));
            int drowsiness = 1;
            int unintended = 1;
            int sudden = 1;
            auto now = chrono::system_clock::now();
            auto timestamp = chrono::system_clock::to_time_t(now);

            //seq++;

            // JSON 형식의 메시지 생성
           // string payload = "{\"seq\": " + to_string(seq) +
           //                 ", \"device_id\": \"" + device_id + "\"" + // 큰따옴표 추가
          //                  ", \"handle\": " + to_string(handle) + 
          //                 ", \"stop\": " + to_string(stop) + 
          //                  ", \"accel\": " + to_string(accel) + 
          //                  ", \"timestamp\": " + to_string(timestamp) + "}";

            string payload = "{\"device_id\": \"" + device_id + "\"" + // 큰따옴표 추가
                            ", \"drowsiness\": " + to_string(drowsiness) + 
                            ", \"unintended\": " + to_string(unintended) + 
                            ", \"sudden\": " + to_string(sudden) + 
                            ", \"timestamp\": " + to_string(timestamp) + "}";

            mqtt::message_ptr pubmsg = mqtt::make_message(TOPIC, payload);
            pubmsg->set_qos(1); // QoS 설정

            client.publish(pubmsg)->wait_for(std::chrono::seconds(10)); // 메시지 전송
            cout << "Sent: " << payload << endl;

            this_thread::sleep_for(chrono::seconds(5)); // 5초 대기
        }
    } catch (const mqtt::exception &e) {
        cerr << "Error: " << e.what() << endl;
        return 1;
    }

    // 연결 종료
    client.disconnect()->wait();
    cout << "Disconnected." << endl;

    return 0;
}


