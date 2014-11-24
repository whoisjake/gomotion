# gomotion

Gomotion is a framework that uses the [WebSocket](https://en.wikipedia.org/wiki/WebSocket) protocol to communicate with a [LeapMotion](https://www.leapmotion.com/) device.

[![baby-gopher](https://raw.github.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)

## Features
* Dead simple to use.
* Frame collection happens concurrently
* JSON comes back parsed and placed into structs for handling.

## Docs

http://godoc.org/github.com/whoisjake/gomotion

## To use

In your $GOPATH:

```bash
$ cd $GOPATH
$ go get github.com/whoisjake/gomotion
```

And then: import "github.com/whoisjake/gomotion"

## Example

```go
package main

import (
	"github.com/whoisjake/gomotion"
	"log"
	"runtime"
)

func main() {
	// Get a device.
	runtime.GOMAXPROCS(runtime.NumCPU())
	device, err := gomotion.GetDevice("ws://127.0.0.1:6437/v3.json")
	if err != nil {
		log.Fatal(err)
	}
	device.Listen()
	defer device.Close()
	for frame := range device.Pipe {
		log.Printf("%+v\n", frame)
	}
}
```
