package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	var mqttServer = flag.String("server", "test.mosquitto.org", "MQTT Server to connect to")
	var mqttPort = flag.Int("port", 1883, "MQTT Server Port")
	var mqttClientId = flag.String("client", "testgoid", "MQTT Client Identifier")
	flag.Parse()

	log.SetOutput(os.Stderr)
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://" + *mqttServer + ":" + strconv.Itoa(*mqttPort))
	opts.SetClientID(*mqttClientId)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := client.Publish("/trackz-autotester/stations", 0, false, text)
		token.Wait()
	}

	client.Disconnect(250)
}
