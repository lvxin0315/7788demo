package config

import "os"

func SetProxyEnv() {
	os.Setenv("http_proxy", "http://127.0.0.1:1087")
	os.Setenv("https_proxy", "http://127.0.0.1:1087")
	os.Setenv("ALL_PROXY", "socks5://127.0.0.1:1080")
}
