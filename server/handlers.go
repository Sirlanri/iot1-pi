package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kataras/iris/v12"
	"github.com/sirlanri/iot1-pi/config"
	"github.com/sirlanri/iot1-pi/log"
)

//NotFound -handler 前端请求地址错误，调用此handler处理
func NotFound(ctx iris.Context) {
	fmt.Println("404-找不到此路由/路径:", ctx.RequestPath(true))
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

	con.WriteString("pi4B: data confirmed")
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
	//log.Log.Errorln("url:", urlPath)

	res, err := http.Get(urlPath)
	if err != nil {
		log.Log.Warnln("向云端发送get请求失败 ", err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Log.Debugln("请求完成", string(body))
}
