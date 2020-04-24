package config

import "github.com/micro/go-micro/v2/config"

var Cfg config.Config

func Init() {
	Cfg, _ = config.NewConfig()
	config.LoadFile("config/config.json")
}
