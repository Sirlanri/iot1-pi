package config

var Dev = true

func BaseurlConf() string {
	if Dev {
		return "http://192.168.1.150:8090/iot1/sensor"
	} else {
		return "https://api.ri-co.cn/iot1/sensor"
	}
}
