package hardware

import (
	"fmt"
	"time"

	"github.com/warthog618/gpiod"
	"github.com/warthog618/gpiod/device/rpi"
)

var (
	Chip *gpiod.Chip //手动控制GPIO
	//Line *gpiod.Line
)

func init() {
	var err error
	Chip, err = gpiod.NewChip("gpiochip0")
	if err != nil {
		fmt.Println("GPIO 初始化芯片组失败 ", err.Error())
		return
	}

}

//Led 操作led： on off blink
func LedSwich(code string) {
	line, err := Chip.RequestLine(rpi.GPIO4, gpiod.AsOutput(0))
	if err != nil {
		fmt.Println("初始化GPIO4出错 ", err.Error())
		return
	}

	defer line.Close()

	if code == "on" {
		line.SetValue(0)
		time.Sleep(time.Second * 2)
		return
	}
	if code == "off" {
		line.SetValue(1)
		time.Sleep(time.Second * 2)
		return
	}
	if code == "blink" {
		for i := 0; i < 5; i++ {
			line.SetValue(0)
			time.Sleep(time.Millisecond * 200)
			line.SetValue(1)
			time.Sleep(time.Millisecond * 200)
		}
		return
	}

}
