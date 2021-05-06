package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-pi/config"
	"github.com/sirlanri/iot1-pi/log"
)

//NotFound -handler 前端请求地址错误，调用此handler处理
func NotFound(ctx iris.Context) {
	log.Log.Warn("404-找不到此路由/路径:", ctx.FullRequestURI())
	ctx.StatusCode(404)
	ctx.WriteString("路由/请求地址错误")
}

//ResEsp -handler 从esp接收温湿度数据
func ResEsp(con iris.Context) {
	humi := con.URLParam("humi")
	temp := con.URLParam("temp")
	light := con.URLParam("light")
	rain := con.URLParam("rain")
	water := con.URLParam("water")
	raininc := con.URLParam("rainincrease")
	log.Log.Debugf("接收到Esp传入 温度：%s,潮湿度%s，光强度%s，下雨%s,水量%s,雨增量%s",
		temp, humi, light, rain, water, raininc)

	//go HumiTempAli(humi, temp)
	go PostAliTemps()

	con.WriteString("pi4B: data confirmed")
}

//-handler 将温度*3发送至云服务器
func TempsAli() {
	temp1 := Temp1
	var temp2, temp3 string

	//判断数据是否空
	if Temp2 == "" {
		temp2 = "0"
	}
	if Temp3 == "" {
		temp3 = "0"
	}

	postData := map[string]interface{}{
		"temp1": temp1,
		"temp2": temp2,
		"temp3": temp3,
	}

	byteData, _ := json.Marshal(postData)
	server := config.BaseurlConf() + "/temps"
	resq, err := http.Post(server, "application/json", bytes.NewBuffer(byteData))
	if err != nil {
		log.Log.Warn("发送温度至服务器失败", err.Error())
		return
	}
	defer resq.Body.Close()
	body, _ := ioutil.ReadAll(resq.Body)
	log.Log.Debug("发送温度至阿里云完成 ", string(body))
}

//-handler 将潮湿度*3发送至云服务器
func HumisAli() {
	humi1 := Temp1
	var humi2, humi3 string

	//判断数据是否空
	if Temp2 == "" {
		humi2 = "0"
	}
	if Temp3 == "" {
		humi3 = "0"
	}

	postData := map[string]interface{}{
		"humi1": humi1,
		"humi2": humi2,
		"humi3": humi3,
	}

	byteData, _ := json.Marshal(postData)
	server := config.BaseurlConf() + "/humis"
	resq, err := http.Post(server, "application/json", bytes.NewBuffer(byteData))
	if err != nil {
		log.Log.Warn("发送潮湿度至服务器失败", err.Error())
		return
	}
	defer resq.Body.Close()
	body, _ := ioutil.ReadAll(resq.Body)
	log.Log.Debug("发送潮湿至服务器完成 ", string(body))
}

//HumiTempAli 将温湿度数据上传到阿里云
func HumiTempAli(humi, temp string) {
	params := url.Values{}
	serverUrl, err := url.Parse(config.BaseurlConf() + "/humitemp?")
	if err != nil {
		log.Log.Errorln("上传温湿度初始化失败 ", err.Error())
		return
	}
	params.Set("humi", humi)
	params.Set("temp", temp)
	urlPath := serverUrl.String() + params.Encode()

	res, err := http.Get(urlPath)
	if err != nil {
		log.Log.Warnln("向云端发送get请求失败 ", err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Log.Debugln("请求完成", string(body))
}

//上传温湿json到server
func PostAliTemps() {
	server := config.BaseurlConf() + "/temps"
	postData := make(map[string]interface{})
	postData["temp1"] = "12.23"
	postData["temp2"] = "12.24"
	postData["temp3"] = "12.25"
	byteData, err := json.Marshal(postData)
	if err != nil {
		log.Log.Warn("post温度至服务器 json化数据失败", err.Error())
		return
	}
	resq, err := http.Post(server, "application/json", bytes.NewBuffer(byteData))
	if err != nil {
		log.Log.Warn("发送温度至服务器失败", err.Error())
		return
	}
	defer resq.Body.Close()
	body, _ := ioutil.ReadAll(resq.Body)
	log.Log.Debug(string(body))
}
