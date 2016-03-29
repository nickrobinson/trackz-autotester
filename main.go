package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"flag"
	"log"
	"os"
	"strconv"
	"time"
	"io"
	"io/ioutil"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	var mqttServer = flag.String("server", "test.mosquitto.org", "MQTT Server to connect to")
	var mqttPort = flag.Int("port", 1883, "MQTT Server Port")
	var mqttClientId = flag.String("client", "testgoid", "MQTT Client Identifier")
	flag.Parse()

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://" + *mqttServer + ":" + strconv.Itoa(*mqttPort))
	opts.SetClientID(*mqttClientId)

	client := MQTT.NewClient(opts)
	Info.Println("Connecting to MQTT server")
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//Keep publishing messages to server continuously
	for {
		Info.Print("Sending test message")
		text := "test message"
		token := client.Publish("/trackz-autotester/stations", 0, false, text)
		token.Wait()
		time.Sleep(5 * time.Second)
	}
}
