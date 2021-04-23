package main

import (
	"sync"

	"github.com/sirlanri/iot1-pi/output"
)

func main() {
	wait := sync.WaitGroup{}
	wait.Add(1)
	output.LedSwich("on")
	wait.Wait()
	return
}
