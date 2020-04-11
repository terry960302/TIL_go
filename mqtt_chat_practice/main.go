package main

import (
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// 디폴트 메세지 핸들러(메시지가 어떤 방식으로 오고 가는지)
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Topic : %s\n", msg.Topic())
	fmt.Printf("Message : %s\n", msg.Payload())
}

func main() {
	// 브로커 설정
	opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt.eclipse.org:1883")
	opts.SetClientID("go-simple")
	opts.SetDefaultPublishHandler(f)

	// 클라이언트
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// go-mqtt/sample 이라는 Topic을 구독
	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// go-mqtt/sample라는 Topic으로 5개의 메세지를 발행
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}

	// 3초 후에 발행을 멈춤.
	time.Sleep(3 * time.Second)

	// go-mqtt/sample의  발행 멈춤
	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}
