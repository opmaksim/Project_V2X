#include <iostream>
#include <mqtt/async_client.h>
#include <nlohmann/json.hpp>
#include <string>
#include <thread>

#define ADDRESS "tcp://10.10.14.34:1883" // MQTT 브로커 주소
#define CLIENT_ID "IotCoreSubscriber"
#define TOPIC "project/v2x/iot-core/car-controller"

using namespace std;
using json = nlohmann::json;

class Callback : public virtual mqtt::callback {
public:
    void message_arrived(mqtt::const_message_ptr msg) override {
        try {
            // 수신한 메시지를 JSON 객체로 파싱
            string payload = msg->to_string();
            json data = json::parse(payload);

            // JSON 데이터에서 필요한 필드를 추출
        //     int seq = data["seq"];
        //     string device_id = data["device_id"];
        //     int handle = data["handle"];
        //     int stop = data["stop"];
        //     int accel = data["accel"];
        //     time_t timestamp = data["timestamp"];

        //     // 데이터 출력
        //     cout << "Message arrived!" << endl;
        //     cout << "Seq: " << seq << endl;
        //     cout << "Device ID: " << device_id << endl;
        //     cout << "Handle: " << handle << endl;
        //     cout << "Stop: " << stop << endl;
        //     cout << "Accel: " << accel << endl;
        //     cout << "Timestamp: " << timestamp << endl;
        //     //cout << "Timestamp: " << timestamp << " (" << ctime(&timestamp) << ")" << endl;
        // } catch (const json::exception &e) {
        //     cerr << "JSON Parsing Error: " << e.what() << endl;
        // }
            int stop_flag = data["stop"];
            cout << "Flage of stop: " << stop_flag << '\n';
        } catch (const json::exception &e) {
                cerr << "JSON Parsing Error: " << e.what() << endl;
        }
    }
};

int main() {
    mqtt::async_client client(ADDRESS, CLIENT_ID);
    mqtt::connect_options connOpts;
    connOpts.set_clean_session(true);

    // Callback 객체 설정
    Callback cb;
    client.set_callback(cb);

    try {
        // 브로커 연결
        cout << "Connecting to the MQTT broker at " << ADDRESS << "..." << endl;
        client.connect(connOpts)->wait();
        cout << "Connected successfully." << endl;

        // 구독 시작
        cout << "Subscribing to topic: " << TOPIC << endl;
        client.subscribe(TOPIC, 1)->wait();
        cout << "Subscribed successfully." << endl;

        // 메시지를 계속 수신
        this_thread::sleep_for(chrono::minutes(10));

    } catch (const mqtt::exception &e) {
        cerr << "Error: " << e.what() << endl;
        return 1;
    }

    // 연결 종료
    client.disconnect()->wait();
    cout << "Disconnected." << endl;

    return 0;
}
