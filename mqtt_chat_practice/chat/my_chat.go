package main

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// https://github.com/yuichi1004/mqtt-chat/blob/master/chat.go

type Chat struct {
	Rooms      map[string]*Room
	MqttClient *client.Client
}

type Room struct {
	Parent   *Chat
	Name     string
	Channels *list.List
}

func NewChat() (*Chat, error) {
	cli := client.New(&client.Options{
		ErrorHandler: func(err error) {
			log.Println(err)
		},
	})
	chat := Chat{MqttClient: cli, Rooms: make(map[string]*Room)}

	u, _ := uuid.NewV4()
	err := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  getMqttMaster() + ":1883",
		ClientID: []byte(u.String()),
	})
	if err != nil {
		return nil, err
	}

	err = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte("rooms/#"),
				QoS:         mqtt.QoS0,
				Handler: func(topicName, message []byte) {
					s := strings.Split(string(topicName), "/")
					if len(s) < 2 {
						log.Printf("Unexpected topic: %s", string(topicName))
						return
					}
					room := chat.GetRoom(s[1])
					var msg Message
					json.Unmarshal(message, &msg)
					for e := room.Channels.Front(); e != nil; e = e.Next() {
						ch := e.Value.(chan Message)
						ch <- msg
					}
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return &chat, nil
}

func (c *Chat) GetRoom(name string) *Room {
	mqttPath := "rooms/" + name
	_, ok := c.Rooms[mqttPath]
	if !ok {
		r := Room{Parent: c, Name: name, Channels: list.New()}
		c.Rooms[mqttPath] = &r
	}
	return c.Rooms[mqttPath]
}

func (r *Room) Subscribe(ch chan Message) {
	r.Channels.PushBack(ch)
}

func (r *Room) Unsubscribe(ch chan Message) {
	for e := r.Channels.Front(); e != nil; e = e.Next() {
		if e.Value == ch {
			r.Channels.Remove(e)
		}
	}
	defer close(ch)
}

func (r *Room) Post(msg Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		msg := fmt.Sprintf("Failed to marshal json: %v", err)
		return errors.New(msg)
	}
	cli := r.Parent.MqttClient
	return cli.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte(fmt.Sprintf("rooms/%s", r.Name)),
		Message:   data,
	})
}
