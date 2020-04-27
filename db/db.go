package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"teamup/config"

	 _ "github.com/jinzhu/gorm/dialects/mysql"
)

var Conn *gorm.DB

func init() {
	ds := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Cfg.Get("db", "username").String(""),
		config.Cfg.Get("db", "password").String(""),
		config.Cfg.Get("db", "host").String(""),
		config.Cfg.Get("db", "database").String(""),
	)

	_db, err := gorm.Open("mysql", ds)
	Conn = _db

	if err != nil {
		log.Fatal("unable to connect to mysql: ", err)
	}
}