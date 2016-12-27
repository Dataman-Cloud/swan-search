package config

import "net"

func DefaultConfig() Config {
	ip := net.ParseIP("0.0.0.0").String()
	config := Config{
		Swans: []Swan{
			Swan{
				Ip:         "172.28.128.4",
				Port:       "9999",
				Scheme:     "http",
				ApiVersion: "v_beta",
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
	Ip         string
	Port       string
	Scheme     string
	ApiVersion string
}
