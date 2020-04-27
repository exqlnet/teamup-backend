package config

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/memory"
	"log"
	"os"
)

var Cfg config.Config

func init() {
	Cfg  = config.DefaultConfig
	var err error

	var sc source.Source
	var cfgData []byte

	switch os.Getenv("GO_RUN_MODE") {
	case "product":
		cfgData, err = Asset("config.json")
		log.Println("Running at product mode")
	case "dev":
		cfgData, err = Asset("config-dev.json")
		log.Println("Running at dev mode")
	case "test":
		cfgData, err = Asset("config-test.json")
		log.Println("Running at test mode")
	default:
		cfgData, err = Asset("config-dev.json")
		log.Println("Running at dev mode")
	}

	sc = memory.NewSource(memory.WithJSON(cfgData))
	err = config.Load(sc)

	if err != nil {
		log.Fatal(err)
	}
}
