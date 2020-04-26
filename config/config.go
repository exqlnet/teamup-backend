package config

import (
	"github.com/micro/go-micro/v2/config"
	"log"
	"os"
)

var Cfg config.Config

func init() {
	Cfg, _ = config.NewConfig()
	if os.Getenv("GO_RUN_MODE") == "product" {
		config.LoadFile("config/config.json")
		log.Println("Running at product mode")
	} else {
		config.LoadFile("config/config-dev.json")
		log.Println("Running at dev mode")
	}

}
