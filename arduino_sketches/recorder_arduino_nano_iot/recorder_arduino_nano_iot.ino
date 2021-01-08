#include <WiFiNINA.h>
#include <ArduinoHttpClient.h>
#include <Arduino_LSM6DS3.h>

char ssid[] = "YOUR_WFIFI_SSID";
char pass[] = "YOUR_WIFI_PASSWORD"; 

int status = WL_IDLE_STATUS;

WiFiClient wifiClient;

const char address[] = "IP.OF.YOUR.SERVER";
int port = 9000;
WebSocketClient client = WebSocketClient(wifiClient, address, port);

void setup() {
  Serial.begin(9600);
  while (!Serial);
  
  // Wifi connect
   if (WiFi.status() == WL_NO_MODULE) {
    Serial.println("No Wifi module found."); 
    // don't continue
    while (true);
  }

  Serial.println("Connecting to Wifi network...");
  status = WiFi.begin(ssid, pass);
  if ( status != WL_CONNECTED) {
    Serial.println("Couldn't get a wifi connection");
    while(true);
  } else {
    Serial.println("Wifi connected!");
  }

  // setup accelerometer and gyro n
  if (!IMU.begin()) {
    Serial.println("Failed to initialize IMU!");
    while (true);
  }

}

void loop() {
  //while(true);
  
  Serial.println("starting WebSocket client");
  client.begin("/record");

  while (client.connected()) {
    float aX, aY, aZ, gX, gY, gZ;
   
    if (IMU.accelerationAvailable()) {
      IMU.readAcceleration(aX, aY, aZ);
    }
  
    if (IMU.gyroscopeAvailable()) {
      IMU.readGyroscope(gX, gY, gZ);
    }
  
    String data = String(aX,2) + "," + String(aY,2) + "," + String(aZ,2) + "," + String(gX,2) + "," + String(gY,2) + "," + String(gZ,2);
    Serial.println(data);
    
    // send data
    client.beginMessage(TYPE_TEXT);
    client.print(data);
    client.endMessage();

    // check if a message is available to be received
    /*int messageSize = client.parseMessage();

    if (messageSize > 0) {
      Serial.println("Received a message:");
      Serial.println(client.readString());
    }*/
    delay(300);
  }

  Serial.println("disconnected");
}
