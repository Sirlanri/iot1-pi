package main

import (
	"github.com/sirlanri/iot1-pi/output"
	"github.com/sirlanri/iot1-pi/server"
)

func main() {
	//sensor.SendHT()
	output.LedSwich("on")
	server.IrisInit()

	return
}
