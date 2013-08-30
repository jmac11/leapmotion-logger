package main

import (
	"github.com/whoisjake/gomotion"
	"log"
	"runtime"
)

func main() {
	// Get a device.
	runtime.GOMAXPROCS(runtime.NumCPU())
	device := gomotion.GetDevice("ws://127.0.0.1:6437/v3.json")
	device.Listen()
	defer device.Close()
	for frame := range device.Pipe {
		log.Printf("%+v\n", frame)
	}
}
