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
table teamup.code_activity_status {

    [0 ]  code                      INT                  nullable: true len: 0
    [1 ]  name                      VARCHAR              nullable: false len: 0
}

*/

// CodeActivityStatus struct is a row record of the code_activity_status table in the teamup database
type CodeActivityStatus struct {
	Code sql.NullInt64
	Name string
}

// TableName sets the insert table name for this struct type
func (c *CodeActivityStatus) TableName() string {
	return "code_activity_status"
}
