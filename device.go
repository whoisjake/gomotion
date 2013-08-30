package gomotion

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

type LeapMotionDevice struct {
	Pipe       chan *Frame
	Connection *websocket.Conn
}

func GetDevice(url string) *LeapMotionDevice {
	pipe := make(chan *Frame)
	connection, err := websocket.Dial(url, "", "http://localhost")
	if err != nil {
		log.Fatal(err)
	}
	return &LeapMotionDevice{pipe, connection}
}

func (device *LeapMotionDevice) Listen() {
	var config struct {
		enableGestures bool `json:"enableGestures"`
	}
	config.enableGestures = true
	err := websocket.JSON.Send(device.Connection, &config)
	if err != nil {
		log.Fatal(err)
	}
	go device.ListenRead()
}

func (device *LeapMotionDevice) ListenRead() {
	for {
		var frame Frame
		err := websocket.JSON.Receive(device.Connection, &frame)
		if err == nil {
			device.Pipe <- &frame
		} else {
			log.Fatal(err)
		}
	}
}

func (device *LeapMotionDevice) Close() {
	device.Connection.Close()
}
