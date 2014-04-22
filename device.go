// The package gomotion defines a concurrent Go library that can connect to a Leap motion device over a WebSocket conection.
// By default, the LeapMotion exposes a JSON WebSocket that pumps out messages near 30 to 50 fps.
package gomotion

import (
	"code.google.com/p/go.net/websocket"
	"net"
)

// A simple type to carry decode errors along rather than just dieing.
type FrameErr struct {
	Frame *Frame
	Error error
}

// The LeapMotionDevice definition. Connecting to a device will return an instance of this struct.
type LeapMotionDevice struct {
	Pipe       chan FrameErr
	Connection *websocket.Conn
}

// This function acts as a constructor and connector for the gomotion package.
func GetDevice(url string) (*LeapMotionDevice, error) {
	pipe := make(chan FrameErr)
	connection, err := websocket.Dial(url, "", "http://localhost")
	if err != nil {
		return nil, err
	}
	return &LeapMotionDevice{pipe, connection}, nil
}

// This function starts the listening on the WebSocket. By default it enables Gestures on the LeapMotionDevice.
func (device *LeapMotionDevice) Listen() error {
	config := struct {
		enableGestures bool `json:"enableGestures"`
	} { true }
	if err := websocket.JSON.Send(device.Connection, &config); err != nil {
		return err
	}
	go device.listenRead()
	return nil
}

func (device *LeapMotionDevice) listenRead() {
	for {
		var frame FrameErr
		if err := websocket.JSON.Receive(device.Connection, &frame.Frame); err != nil {
			// Slightly way to avoid another variable in LeapMotionDevice. Works for me.
			if _, ok := err.(*net.OpError); ok {
				close(device.Pipe)
				return
			}
			frame.Error = err
		}
		device.Pipe <- frame
	}
	close(device.Pipe)
}

// This function closes the internal WebSocket connection on a LeapMotionDevice
func (device *LeapMotionDevice) Close() {
	device.Connection.Close()
}
