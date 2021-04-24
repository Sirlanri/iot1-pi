package sensor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/d2r2/go-dht"
	"github.com/sirlanri/iot1-pi/config"
)

func SendHT() {
	for {
		temperature, humidity, _ :=
			dht.ReadDHTxx(dht.DHT11, 17, false)
		if temperature != -1 {
			fmt.Println("采集的数据：", temperature, humidity)
			sendToServer(float64(temperature), float64(humidity))
			time.Sleep(time.Second)
		}
	}
}
func sendToServer(temp, humi float64) {
	params := url.Values{}
	serverUrl, err := url.Parse(config.BaseurlConf() + "/humitemp?")
	if err != nil {
		fmt.Println("上传温湿度初始化失败 ", err.Error())
		return
	}
	params.Set("humi", strconv.FormatFloat(humi, 'f', 2, 32))
	params.Set("temp", strconv.FormatFloat(temp, 'f', 2, 32))
	urlPath := serverUrl.String() + params.Encode()
	fmt.Println("url:", urlPath)

	res, err := http.Post(urlPath, "", nil)
	if err != nil {
		fmt.Println("发送post请求失败 ", err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("post完成", string(body))
}
