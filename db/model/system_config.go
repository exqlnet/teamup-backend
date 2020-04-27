package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

/*
DB Table Details
table teamup.system_config {

    [0 ]  key                       VARCHAR              nullable: false len: 0
    [1 ]  value                     VARCHAR              nullable: false len: 0
}

*/

// SystemConfig struct is a row record of the system_config table in the teamup database
type SystemConfig struct {
	Key   string
	Value string
}

// TableName sets the insert table name for this struct type
func (s *SystemConfig) TableName() string {
	return "system_config"
}
