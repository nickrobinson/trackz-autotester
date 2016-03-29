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
	"github.com/nickrobinson/trackz-puller/stations"
	"encoding/json"
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
		s := stations.Stations{
			[]stations.StationInfo{
				stations.StationInfo{"FooAP", "00:04:F2:80:21:BC", -34},
			},
		}
		b, _ := json.Marshal(s)
		Info.Print("Sending test message")
		text := string(b)
		token := client.Publish("/nickrobi/0001/aps", 0, false, text)
		token.Wait()
		time.Sleep(5 * time.Second)
	}
}
