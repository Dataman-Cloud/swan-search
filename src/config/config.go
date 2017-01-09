package config

import "net"

func DefaultConfig() Config {
	ip := net.ParseIP("0.0.0.0").String()
	config := Config{
		Swans: []Swan{
			Swan{
				Urls: "http://172.28.128.4:9999",
			},
		},
		Ip:     ip,
		Port:   "80",
		Scheme: "http",
	}
	return config
}

type Config struct {
	Swans  []Swan
	Ip     string
	Port   string
	Scheme string
}

type Swan struct {
	Urls string
}
