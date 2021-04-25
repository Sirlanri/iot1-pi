package output

import (
	"fmt"
	"time"

	"github.com/sirlanri/iot1-pi/config"
	"github.com/warthog618/gpiod"
	"github.com/warthog618/gpiod/device/rpi"
)

var (
	Chip *gpiod.Chip //手动控制GPIO
	//Line *gpiod.Line
	Instruction string //控制LED的信标
)

func init() {
	var err error
	Chip, err = gpiod.NewChip("gpiochip0")
	if err != nil {
		fmt.Println("GPIO 初始化芯片组失败 ", err.Error())
		return
	}
	Instruction = "off"

	line, err := Chip.RequestLine(rpi.GPIO4, gpiod.AsOutput(0))

	if err != nil {
		fmt.Println("初始化GPIO4出错 ", err.Error())
		return
	}

	go ledControl(line)

	//line.Close()
}

//Led 操作led： on off blink
func LedSwich(code string) {
	config.Log.Debugln("测试级别")
	config.Log.Warnln("警告级别")
	Instruction = code
}

//LedControl 控制LED灯的实际函数，应放入一个线程中运行
func ledControl(line *gpiod.Line) {
	for {
		if Instruction == "on" {
			line.SetValue(0)
			//time.Sleep(time.Second)
		}
		if Instruction == "off" {
			line.SetValue(1)
			time.Sleep(time.Second)
		}
		if Instruction == "blink" {
			for i := 0; i < 5; i++ {
				line.SetValue(0)
				time.Sleep(time.Millisecond * 200)
				line.SetValue(1)
				time.Sleep(time.Millisecond * 200)
			}
		}

		//fmt.Println("正在执行控制 ", Instruction)
	}

}
