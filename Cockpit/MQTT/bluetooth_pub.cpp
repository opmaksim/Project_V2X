#include <iostream>
#include <string>
#include <thread>
#include <chrono>
#include <mqtt/async_client.h>
#include <stdexcept>
#include <vector>
#include <ctime>
#include <csignal>
#include <unistd.h>
#include <bluetooth/bluetooth.h>
#include <bluetooth/rfcomm.h>

#define ADDRESS "tcp://10.10.14.34:1883"
#define CLIENT_ID "SCM_CAR"
#define TOPIC "project/v2x/iot-core/data-drive-collector"
#define BUF_SIZE 1024

using namespace std;

// Bluetooth 연결 함수 (C++ 스타일로 예외 처리 적용)
int connect_bluetooth(const string& address) {
    int bt_sock;
    struct sockaddr_rc addr = {0};

    bt_sock = socket(AF_BLUETOOTH, SOCK_STREAM, BTPROTO_RFCOMM);
    if (bt_sock == -1) {
        throw runtime_error("Failed to create Bluetooth socket");
    }

    addr.rc_family = AF_BLUETOOTH;
    addr.rc_channel = (uint8_t) 1;
    str2ba(address.c_str(), &addr.rc_bdaddr);

    if (connect(bt_sock, (struct sockaddr*)&addr, sizeof(addr)) == -1) {
        close(bt_sock);
        throw runtime_error("Failed to connect to Bluetooth device");
    }

    cout << "Bluetooth connected successfully." << endl;
    return bt_sock;
}

// Bluetooth 데이터 수신 및 MQTT 전송 루프
void bluetooth_recv_loop(int bt_sock, mqtt::async_client& client) {
    vector<char> buf(BUF_SIZE);

    while (true) {
        int ret = read(bt_sock, buf.data(), BUF_SIZE - 1);
        if (ret > 0) {
            buf[ret] = '\0';  // 문자열 종료
            string msg(buf.data());
            cout << "Received from Bluetooth: " << msg << endl;

            // MQTT로 메시지 전송
            mqtt::message_ptr pubmsg = mqtt::make_message(TOPIC, msg);
            pubmsg->set_qos(1);

            try {
                client.publish(pubmsg)->wait_for(chrono::seconds(10));
                cout << "Sent to MQTT: " << msg << endl;
            } catch (const mqtt::exception& e) {
                cerr << "Failed to send to MQTT: " << e.what() << endl;
            }
        } else if (ret == 0) {
            cout << "Bluetooth connection closed." << endl;
            break;
        } else {
            perror("Bluetooth read");
            break;
        }
    }

    close(bt_sock);
}

int main() {
    mqtt::async_client client(ADDRESS, CLIENT_ID);
    mqtt::connect_options connOpts;
    connOpts.set_clean_session(true);

    try {
        cout << "Connecting to the MQTT broker at " << ADDRESS << "..." << endl;
        client.connect(connOpts)->wait();
        cout << "Connected to MQTT broker." << endl;

        // 블루투스 연결 (C++ 스타일로 주소를 string으로 사용)
        string bt_address = "98:DA:60:07:62:B4";
        int bt_sock = connect_bluetooth(bt_address);

        // 블루투스 데이터 수신 루프 실행 (스레드로 처리)
        thread bluetooth_thread(bluetooth_recv_loop, bt_sock, ref(client));
        bluetooth_thread.join();  // 블루투스 스레드 종료 대기

    } catch (const exception& e) {
        cerr << "Error: " << e.what() << endl;
        return 1;
    }

    client.disconnect()->wait();
    cout << "Disconnected from MQTT broker." << endl;
    return 0;
}