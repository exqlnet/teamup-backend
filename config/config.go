package config

import (
	"github.com/micro/go-micro/v2/config"
	"log"
	"os"
)

var Cfg config.Config
var CfgMap map[string]interface{}

func init() {
	Cfg  = config.DefaultConfig
	cwd, _ := os.Getwd()
	log.Println("Current dir is", cwd)
	cfgDir := cwd + string(os.PathSeparator) + "config" + string(os.PathSeparator)
	var err error
	switch os.Getenv("GO_RUN_MODE") {
	case "product":
		err = config.LoadFile(cfgDir + "config.json")
		log.Println("Running at product mode")
	case "dev":
		err = config.LoadFile(cfgDir + "config-dev.json")
		log.Println("Running at dev mode")
	case "test":
		err = config.LoadFile(cfgDir + "config-test.json")
		log.Println("Running at test mode")
	default:
		err = config.LoadFile(cfgDir + "config-dev.json")
		log.Println("Running at dev mode")
	}

	if err != nil {
		log.Fatal(err)
	}
	CfgMap = Cfg.Map()
}
