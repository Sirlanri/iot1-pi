package server

import (
	"encoding/json"
	"fmt"

	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"github.com/sirlanri/iot1-pi/log"
	"github.com/sirlanri/iot1-pi/output"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("主题: %s\n", msg.Topic())
	fmt.Printf("消息内容: %s\n", msg.Payload())

}

var msgHandler = func(client mqtt.Client, msg mqtt.Message) {
	topic := string(msg.Topic())

	//控制led
	if topic == "iot1/pi/res/led" {
		output.LedSwich(string(msg.Payload()))
	}

}

var c mqtt.Client

//Createid 生成唯一名称
func Createid() string {
	// 创建 UUID v4
	u1 := uuid.Must(uuid.NewV4(), nil)
	id := u1.String()
	return id[:9]
}

//初始化
func init() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt.ri-co.cn:1883").SetClientID("emqx_golang_" + Createid())

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sub()

}

//sub 订阅某个主题的消息 控制硬件的指令
func sub() {
	topic := "iot1/pi/res/#"
	token := c.Subscribe(topic, 2, msgHandler)
	token.Wait()

}

//SendMqtt 通过mqtt发送消息
func SendMqtt(payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Log.Errorln("json打包出错", err.Error())
	}
	go func() {
		token := c.Publish("iot1/server/send", 2, false, data)
		err = token.Error()
		if err != nil {
			log.Log.Errorln("mqtt出错", err.Error())
		}
		token.Wait()
	}()

}

//SendMqttString 通过mqtt发送消息，文本格式
func SendMqttString(payload string) {
	go func() {
		token := c.Publish("iot1/server/send", 2, false, payload)
		err := token.Error()
		if err != nil {
			log.Log.Errorln("mqtt出错", err.Error())
		}
		token.Wait()
	}()
}

//SendMqttInfo 通过mqtt发送报错/消息，文本格式
func SendMqttInfo(payload string) {
	go func() {
		token := c.Publish("iot1/server/info", 2, false, payload)
		err := token.Error()
		if err != nil {
			log.Log.Errorln("mqtt出错", err.Error())
		}
		token.Wait()
	}()
}
