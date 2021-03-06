package cfg

import "flag"

const (
	OK            = "OK"
	HEALTH        = "Health"
	UNHEALTHY     = "Unhealthy"
	VERSION       = "VERSION"
	ERROR_HANDLER = "ERROR"
)

type ServerConfig struct {
	Address string
}

func SetUp() *ServerConfig {
	config := ServerConfig{
		Address: "8080",
	}
	flag.StringVar(&config.Address, "address", "8080", "设置服务器监听地址")
	flag.Parse()
	return &config
}
