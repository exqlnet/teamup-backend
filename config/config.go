package config

import (
	"github.com/micro/go-micro/v2/config"
)

var Cfg config.Config

func init() {
	Cfg, _ = config.NewConfig()
	config.LoadFile("config/config.json")
}
