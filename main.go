package main

import (
	"sync"

	"github.com/sirlanri/iot1-pi/output"
	"github.com/sirlanri/iot1-pi/server"
	_ "github.com/sirlanri/iot1-pi/server"
)

func main() {
	wait := sync.WaitGroup{}
	wait.Add(1) //阻塞准备

	//sensor.SendHT()
	output.LedSwich("on")
	server.IrisInit()
	//阻塞
	wait.Wait()
	return
}
