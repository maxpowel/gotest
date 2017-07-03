package main

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/eclipse/paho.mqtt.golang"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/fatih/color"
)

type MqttConfig struct {
	Hostname string `default:"valor"`
	Port int `default:"PUERTOOO"`
	Topic string
}

type PlayList struct {
	gorm.Model
	Source string
	SourceType string
}

func mqttBootstrap(k *Kernel) {
	mapping := k.config.mapping
	mapping["mqtt"] = &MqttConfig{}

	var baz OnKernelReady = func(k *Kernel){
		color.Green("Evento %v en mqtt")
		conf := k.config.mapping["mqtt"].(*MqttConfig)

		//conf = k.config.mapping["mqtt"]
		// Start mqtt connection
		//opts := mqtt.NewClientOptions().AddBroker("tcp://iot.eclipse.org:1883").SetClientID("gotrivial")
		opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%v:%v", conf.Hostname, conf.Port))

		//opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%v:%v", "a", "b"))
		opts.SetKeepAlive(2 * time.Second)
		var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
			color.Blue("TOPIC: %s\n", msg.Topic())
			newTest := &Respuesta{}
			/*db, err := gorm.Open("mysql", "mqtt:123456@tcp(localhost:3306)/mqtt?charset=utf8&parseTime=true")
			// Migrate the schema
			db.AutoMigrate(&PlayList{})
			if err != nil {
				panic("failed to connect database")
			}
			defer db.Close()*/


			err := proto.Unmarshal(msg.Payload(), newTest)
			if err != nil {
				log.Fatal("unmarshaling error: ", err)
			}
			//db.Create(&PlayList{Source: newTest.Message, SourceType: string(newTest.StatusCode)})
			db2 := k.container.MustGet("database").(*gorm.DB)
			//db := container.MustGet("gorm").(gorm.DB)
			db2.Create(&PlayList{Source: "AAA", SourceType: "BBB"})
			color.Blue("MSG: %s\n", newTest.Message)
		}
		opts.SetDefaultPublishHandler(f)
		opts.SetPingTimeout(1 * time.Second)

		c := mqtt.NewClient(opts)
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			color.Red(token.Error().Error())
			return
		}
		color.Green("MQTT connection established with %v", conf.Hostname)
		if token := c.Subscribe(conf.Topic, 0, nil); token.Wait() && token.Error() != nil {
			color.Red(token.Error().Error())
			return
		}
		color.Green("Subscribed to %v", conf.Topic)

		defer func() {
			color.Green("Disconnecting")
			c.Disconnect(250)
		}()

		defer func() {
			color.Green("Unsubscribing")
			unsubscribeToken := c.Unsubscribe(conf.Topic)
			unsubscribeToken.Wait()
		}()



		test := &Respuesta {
			Message: "tuvieja",
			StatusCode: 24,
		}

		data, err := proto.Marshal(test)

		fmt.Println(data)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		ioutil.WriteFile("mensaje", data, 0644)

		/*b, err := json.Marshal(group)
		if err != nil {
			fmt.Println("error:", err)
		}

		os.Stdout.Write(b)*/





		/*for i := 0; i < 5; i++ {
			text := fmt.Sprintf("this is msg #%d!", i)
			token := c.Publish("go-mqtt/sample", 0, false, text)
			token.Wait()
		}*/

		//time.Sleep(6 * time.Second)
		daemonize()
	}
	k.subscribe(baz)


}
