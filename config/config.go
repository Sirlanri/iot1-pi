package config

var Dev = false

func BaseurlConf() string {
	if Dev {
		return "http://192.168.1.150:8090/iot1/pi"
	} else {
		return "https://api.ri-co.cn/iot1/pi"
	}
}
